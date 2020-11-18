package main

/*
	Demonstrates a simple window.
*/

import(
	"github.com/prospero78/goGui/lib"
	"github.com/prospero78/goGui/lib/window"
)

func main(){
	log:=lib.Log
	win:=window.NewWindow("")
	win.SetTitle("Simple window")
	win.Run()
	log.Infof("Window is close")
}
