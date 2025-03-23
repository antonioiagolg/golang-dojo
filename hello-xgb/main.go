package main

import (
	"fmt"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

func main() {
    X, err := xgb.NewConn()

    if err != nil {
        fmt.Println(err)
        return
    }

    wid, _ := xproto.NewWindowId(X)
    screen := xproto.Setup(X).DefaultScreen(X)

    xproto.CreateWindow(X, screen.RootDepth, wid, screen.Root, 0, 0, 500, 500, 0,
        xproto.WindowClassInputOutput, screen.RootVisual, xproto.CwBackPixel | xproto.CwEventMask,
        []uint32{
            0xffffffff,
            xproto.EventMaskStructureNotify |
            xproto.EventMaskKeyPress |
            xproto.EventMaskKeyRelease})

    xproto.MapWindow(X, wid)

    for {
        ev, xerr := X.WaitForEvent()
        if ev == nil && xerr == nil {
            fmt.Println("Both envent and err are nil, exiting...")
            return
        }
        if ev != nil {
            fmt.Printf("Event: %s\n", ev)
        }
        if xerr != nil {
            fmt.Printf("Err: %s\n", ev)
        }
    }
}
