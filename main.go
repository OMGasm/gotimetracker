package main

import "log/slog"
import "github.com/jezek/xgb"
import "github.com/jezek/xgb/xproto"
import "example.com/gotimetracker/x"

var activeAtom, nameAtom *xproto.InternAtomReply

func main() {
	X, err := x.Init()
	defer X.Close()

	reply, err := X.Prop(X.Atoms.Active.Atom, X.Root())
	if err != nil {
		slog.Error("error?", "err", err)
	}
	windowId := xproto.Window(xgb.Get32(reply.Value))
	slog.Info("Active window id", "id", windowId)

	reply, err = X.Prop(X.Atoms.WindowName.Atom, windowId)
	if err != nil {
		slog.Error("error?", "err", err)
	}
	slog.Info("Active window name", "name", string(reply.Value))
}
