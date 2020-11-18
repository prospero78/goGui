package window

/*
	Пакет предоставляет тип окна. Все отрисовки производятся именно в нём.

	The package provides the window type. All rendering is done in it.
*/

import (
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/zserge/lorca"

	"github.com/prospero78/goGui/lib"
)

const (
	page = `<HTML>
	<HEAD>
		<TITLE>{{title}}</TITLE>
		<META CHARSET="UTF-8"/>
	</HEAD>
	<BODY>
	</BODY>
	</HTML>`
)

var (
	countWin = 0 // count for all window
	block    sync.RWMutex
)

// TWindow -- operation with window
type TWindow struct {
	ui      lorca.UI    // User interface
	numWin  int         // Number win
	chTitle chan string //Chan for setting title window
}

// NewWindow -- returns new *TWindow
func NewWindow(title string) (window *TWindow) {
	defer block.Unlock()
	block.Lock()
	countWin++
	if title == "" {
		title = fmt.Sprintf("win%v", countWin)
	}
	strPage := strings.ReplaceAll(page, "{{title}}", title)
	ui, err := lorca.New("data:text/html,"+url.PathEscape(strPage), "", 480, 360)
	if err != nil {
		lib.Log.Panicf("NewWindow(): PANIC in create window\n\t%v", err)
	}

	window = &TWindow{
		ui:      ui,
		numWin:  countWin,
		chTitle: make(chan string, 5),
	}
	return window
}

// SetTitle -- set title window on demand
func (sf *TWindow) SetTitle(title string) {
	sf.chTitle <- title
}

// Run -- run mainloop of window
func (sf *TWindow) Run() {
	for {
		select {
		case <-sf.ui.Done(): // Close window
			sf.ui.Close()
			return
		case title := <-sf.chTitle: //Change title
			sf.ui.Eval(`document.title="` + title + `";`)
		}
	}
}
