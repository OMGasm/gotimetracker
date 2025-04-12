package main

import (
	"fmt"

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
			fmt.Printf("%s: %02d:%02d:%02d\n", w, h, m, s)
		}
	}
}
