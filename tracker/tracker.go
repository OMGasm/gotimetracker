package tracker

import (
	"log/slog"
	"maps"
	"sync"
	"time"

	"github.com/OMGasm/gotimetracker/x"
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

type Tracker struct {
	interval time.Duration
	timer    *time.Ticker
	wmx      sync.Mutex
	windows  map[string]time.Time
	close    chan bool
	x        *x.X
}

func New(x *x.X) *Tracker {
	t := new(Tracker)
	t.interval = 1 * time.Second
	t.timer = time.NewTicker(1 * time.Second)
	t.windows = make(map[string]time.Time)
	t.close = make(chan bool)
	t.x = x

	go t.loop()
	return t
}

func (self *Tracker) loop() {
	for {
		select {
		case <-self.close:
			break
		case <-self.timer.C:
			wprop, err := self.x.Prop(self.x.Atoms.Active.Atom, self.x.Root())
			if err != nil {
				slog.Error("Error", "window id", err)
				continue
			}
			windowId := xproto.Window(xgb.Get32(wprop.Value))
			wnProp, err := self.x.Prop(self.x.Atoms.WindowName.Atom, windowId)
			if err != nil {
				slog.Error("Error", "window name", err)
			}
			windowName := string(wnProp.Value)
			{
				self.wmx.Lock()
				self.windows[windowName] = self.windows[windowName].Add(self.interval)
				self.wmx.Unlock()
			}
		}
	}
}

func (self *Tracker) Entries() map[string]time.Time {
	self.wmx.Lock()
	windows := maps.Clone(self.windows)
	self.wmx.Unlock()

	return windows
}

func (self *Tracker) Close() {
	self.timer.Stop()
	self.close <- true
}

func (self *Tracker) Start() {
	self.timer.Reset(self.interval)
}

func (self *Tracker) Stop() {
	self.timer.Stop()
}
