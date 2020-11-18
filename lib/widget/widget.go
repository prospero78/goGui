package widget

/*
	-ru-
	Корневой тип виджета для всех виджетов. Содержит реализацию всех необходимых методов.

	-en-
	Root widget type for all widgets. Contains the implementation of all required methods.
*/

import (
	"github.com/prospero78/goGui/lib/log"
	"github.com/prospero78/goGui/lib/types"
	"github.com/prospero78/goGui/lib/widget/counterwidgets"
	"github.com/prospero78/goGui/lib/widget/widgetid"
)

// TWidget -- operations with root-widget
type TWidget struct {
	*widgetid.TWidgetID
	*counterwidgets.TCounterWidgets
	parent types.IParent
	log    *log.TLog
}

// NewWidget -- return new *TWidget
func NewWidget(parent types.IParent) (wd *TWidget) {
	lg := log.NewLog()
	lg.SetLevel(log.DEBUG)
	lg.SetPrefix("TWidget")
	counterwidgets.CounterWidgets.Inc()
	wd = &TWidget{
		TWidgetID:       widgetid.NewWidgetID(counterwidgets.CounterWidgets.Get()),
		TCounterWidgets: counterwidgets.CounterWidgets,
		parent:          parent,
		log:             lg,
	}
	return wd
}

// AddWidget -- add self in parent widget. Need individual reallisation!!
func (sf *TWidget) AddWidget(widget types.IWidget) {
	sf.log.Panicf("AddWidget(): need individual realisation for all widget\n")
}

// GetParent -- return parent of widget
func (sf *TWidget) GetParent() types.IParent {
	return sf.parent
}
