package main

import (
	"fmt"
	"time"

	"example.com/gotimetracker/tracker"
	"example.com/gotimetracker/x"
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
