package main

/*
	Demonstrates a simple window.
*/

import(
	"github.com/prospero78/goGui/lib"
)

func main(){
	log:=lib.Log
	win, err:=lib.NewWin("Simple window")
	if err!=nil{
		log.Panicf("panic in create simple window\n\t%v", err)
	}
}
