package x

import (
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
	"log/slog"
)

type X struct {
	conn          *xgb.Conn
	setup         *xproto.SetupInfo
	defaultScreen *xproto.ScreenInfo
	root          xproto.Window
	Atoms         struct {
		Active     *xproto.InternAtomReply
		WindowName *xproto.InternAtomReply
	}
}

func Init() (*X, error) {
	x, err := xgb.NewConn()
	if err != nil {
		slog.Error("Fatal error", "Could not init X connection", err)
		panic(err)
	}
	self := &X{conn: x}

	intern := func(name string) *xproto.InternAtomReply {
		atom, err := xproto.InternAtom(x, true, uint16(len(name)), name).Reply()
		if err != nil {
			slog.Error("Fatal error", "Could not intern atom", err)
			panic(err)
		}
		return atom
	}

	self.Atoms.Active = intern("_NET_ACTIVE_WINDOW")
	self.Atoms.WindowName = intern("_NET_WM_NAME")

	self.setup = xproto.Setup(x)
	self.defaultScreen = self.setup.DefaultScreen(x)
	self.root = self.defaultScreen.Root

	return self, err
}

func (self *X) Prop(propAtom xproto.Atom, from xproto.Window) (*xproto.GetPropertyReply, error) {
	reply, err := xproto.GetProperty(self.conn, false, from, propAtom, xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
	return reply, err
}

func (self *X) Close() {
	self.conn.Close()
}

func (self *X) Root() xproto.Window { return self.root }
