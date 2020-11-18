package main

/*
	Demonstrates a simple window.
*/

import (
	"time"

	"github.com/prospero78/goGui/lib"
	"github.com/prospero78/goGui/lib/window"
)

func main() {
	log := lib.Log
	win := window.NewWindow("")

	go win.Run()

	time.Sleep(time.Second * 2)
	win.SetTitle("1) Simple window")

	time.Sleep(time.Second * 2)
	win.SetTitle("2) Resize window")
	win.SetSize(480, 180)

	time.Sleep(time.Second * 2)
	win.SetTitle("3) Set background window")
	win.SetColorBg("#404040")

	time.Sleep(time.Second * 2)
	win.SetTitle("4) Set background image")
	win.SetSize(640, 480)
	win.SetImageBg("https://img3.goodfon.ru/original/640x480/8/53/okean-more-voda-skaly-derevya.jpg")


	time.Sleep(time.Second * 2)
	win.SetTitle("4) Set fixed and set size")
	win.SetFixed()
	win.SetSize(640, 480)

	time.Sleep(time.Second * 2)
	win.SetTitle("4) Set unfixed and set size")
	win.SetUnfixed()
	win.SetSize(645, 485)

	time.Sleep(time.Second * 2)
	log.Infof("Window is closed")
	win.SetTitle("4) Close window")
	time.Sleep(time.Millisecond * 700)
	win.Close()
}
