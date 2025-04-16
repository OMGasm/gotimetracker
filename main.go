package main

import (
	"fmt"
	"time"

	"github.com/OMGasm/gotimetracker/tracker"
	"github.com/OMGasm/gotimetracker/x"
	"github.com/jezek/xgb/xproto"
)

var activeAtom, nameAtom *xproto.InternAtomReply

func main() {
	X, err := x.Init()
	if err != nil {
		return
	}
	defer X.Close()
	tracker := tracker.New(X)
	tracker.Start()

	for {
		fmt.Print("\f")
		foo := tracker.Entries()

		for w, t := range foo {
			h, m, s := t.Clock()
			fmt.Printf("[%02d:%02d:%02d] %s\n", h, m, s, w)
		}
		time.Sleep(1 * time.Second)
	}
}
