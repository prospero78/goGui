package window

/*
	Пакет предоставляет тип окна. Все отрисовки производятся именно в нём.

	The package provides the window type. All rendering is done in it.
*/

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/zserge/lorca"

	"github.com/prospero78/goGui/lib/css"
	"github.com/prospero78/goGui/lib/log"
	"github.com/prospero78/goGui/lib/size"
	"github.com/prospero78/goGui/lib/types"
	"github.com/prospero78/goGui/lib/widget"
)

const (
	defaultSizeX = 480
	defaultSizeY = 360
	page         = `<HTML>
	<HEAD>
		<TITLE>{{title}}</TITLE>
		<META CHARSET="UTF-8"/>
		<STYLE>{{style}}</STYLE>
	</HEAD>
	<BODY>
	</BODY>
	</HTML>`
	sizeFixed = iota
	sizeUnfixed
)

// TWindow -- operation with window
type TWindow struct {
	*widget.TWidget
	ui          lorca.UI                          // User interface
	chTitle     chan string                       //Chan for setting title window
	chSize      chan *size.TSize                  // Chan for set size window
	size        *size.TSize                       // Size of window
	chClose     chan int                          // Chan for close window
	chColorBg   chan string                       // Set background color
	chImageBg   chan string                       // Set background image
	chSizeFixed chan int                          // Chan for set fixed/unfixed size window
	chWidgetAdd chan types.IWidget                // Chan for add widget in window
	log         *log.TLog                         // Local log for out info, error , etc.
	poolWidget  map[types.AWidgetID]types.IWidget //dictionary widgets in window
}

// NewWindow -- returns new *TWindow
func NewWindow(title string) (window *TWindow) {
	widget := widget.NewWidget(nil)
	if title == "" {
		title = fmt.Sprintf("win%v", widget.GetWidgetID())
	}
	lg := log.NewLog()
	lg.SetPrefix(fmt.Sprintf("%v TWindow", title))
	lg.SetLevel(log.DEBUG)

	strPage := strings.ReplaceAll(page, "{{title}}", title)
	strPage = strings.ReplaceAll(page, "{{style}}", css.GetBody())

	ui, err := lorca.New("data:text/html,"+url.PathEscape(strPage), "", defaultSizeX, defaultSizeY)
	if err != nil {
		lg.Panicf("NewWindow(): PANIC in create window\n\t%v", err)
	}

	window = &TWindow{
		TWidget:     widget,
		ui:          ui,
		chTitle:     make(chan string, 5),
		chSize:      make(chan *size.TSize, 5),
		chClose:     make(chan int, 5),
		chColorBg:   make(chan string, 5),
		chImageBg:   make(chan string, 5),
		chSizeFixed: make(chan int, 5),
		chWidgetAdd: make(chan types.IWidget, 5),
		size:        size.NewSize(defaultSizeX, defaultSizeY),
		log:         lg,
		poolWidget:  make(map[types.AWidgetID]types.IWidget),
	}
	return window
}

// AddWidget -- add widget in internal space
func (sf *TWindow) AddWidget(widget types.IWidget) {
	if widget == nil {
		sf.log.Errorf("AddWidget(): widget==nil\n")
		return
	}
	sf.chWidgetAdd <- widget
}

// SetFixed -- set fixed size window
func (sf *TWindow) SetFixed() {
	sf.chSizeFixed <- sizeFixed
	time.Sleep(time.Millisecond * 10)
}

// SetUnfixed -- set unfixed size window
func (sf *TWindow) SetUnfixed() {
	sf.chSizeFixed <- sizeUnfixed
	time.Sleep(time.Millisecond * 10)
}

// SetColorBg -- set background color in window
func (sf *TWindow) SetColorBg(color string) {
	sf.chColorBg <- color
}

// SetImageBg -- set background image in window
func (sf *TWindow) SetImageBg(image string) {
	sf.chImageBg <- image
}

// Close -- close window on demand
func (sf *TWindow) Close() {
	sf.chClose <- 1
	time.Sleep(time.Millisecond * 50)
}

// SetSize -- change size window to absolute size (x; y)
func (sf *TWindow) SetSize(sizeX size.ASizeX, sizeY size.ASizeY) {
	if sizeX < 0 {
		sizeX = 0
	}
	if sizeY < 0 {
		sizeY = 0
	}
	sf.chSize <- size.NewSize(sizeX, sizeY)
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
		case size := <-sf.chSize: //Change size
			if sf.size.IsFixed() {
				sf.log.Errorf("Run(): size window is fixed")
				continue
			}
			sf.size = size
			strX := fmt.Sprint(size.GetX())
			strY := fmt.Sprint(size.GetY())
			cmd := fmt.Sprintf(`window.resizeTo(%v, %v);`, strX, strY)
			sf.ui.Eval(cmd)
		case <-sf.chClose: // Close window
			sf.ui.Eval(`window.close();`)
		case color := <-sf.chColorBg: //Set background color
			sf.ui.Eval(`document.body.style.background ="` + color + `"`)
		case image := <-sf.chImageBg: // Set background image
			sf.ui.Eval(`document.body.style.background ="url(` + image + `)"`)
		case fixed := <-sf.chSizeFixed: // Change fixed size window
			switch fixed {
			case sizeFixed:
				sf.size.SetFixed()
			case sizeUnfixed:
				sf.size.ResetFixed()
			}
		case widget := <-sf.chWidgetAdd: // Add widget in window
			sf.addWidget(widget)
		}
	}
}

func (sf *TWindow) addWidget(widget types.IWidget) {
	sf.poolWidget[widget.GetWidgetID()] = widget
	strID := fmt.Sprint(widget.GetWidgetID())
	sf.ui.Eval(`
			let div = document.createElement('DIV');
			div.id = '` + strID + `';
			document.body.append(div);`)
	style := widget.GetStyle()
	for side, st := range style {
		color := st["Color"]
		cmd:=`document.getElementById('` + strID + `').style.` + side + `Color='` + color + `';`
		sf.ui.Eval(cmd)

		width := st["Width"]
		cmd=`document.getElementById('` + strID + `').style.` + side + `Width='` + width + `';`
		sf.ui.Eval(cmd)

		style := st["Style"]
		cmd=`document.getElementById('` + strID + `').style.` + side + `Style='` + style + `';`
		sf.ui.Eval(cmd)
	}
}
