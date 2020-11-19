package frame

import (
	"github.com/prospero78/goGui/lib/attr/linestyle"
	"github.com/prospero78/goGui/lib/border"
	"github.com/prospero78/goGui/lib/log"
	"github.com/prospero78/goGui/lib/types"
	"github.com/prospero78/goGui/lib/widget"
)

/*
	-ru-
	Пакет предоставляет тип простого фрейма. Его задача -- группировать в себе элементы.

	-en-
	The package provides a simple frame type. Its task is to group elements in itself.
*/

// TFrame -- frame operations
type TFrame struct {
	*widget.TWidget
	border *border.TBorder
}

// NewFrame -- return new *TFrame
func NewFrame(parent types.IParent) (frame *TFrame) {
	lg := log.NewLog()
	lg.SetLevel(log.DEBUG)
	lg.SetPrefix("TFrame")
	{ // Precondition
		if parent == nil {
			lg.Panicf("NewFrame(): parent==nil\n")
		}
	}
	frame = &TFrame{
		TWidget: widget.NewWidget(parent),
		border:  border.NewBorder("#404040", 2, linestyle.Double),
	}
	return frame
}

// Border -- return saved border
func (sf *TFrame) Border() *border.TBorder {
	return sf.border
}

// SetThickness -- set borders thickness
func (sf *TFrame) SetThickness(thickNes types.AThickness) {
	sf.border.SetThickness(thickNes)
	// update visual style
	sf.TWidget.GetParent().AddWidget(sf)
}

// GetHTML -- return HTML-represent TFrame
func (sf *TFrame) GetHTML() types.AHtml {
	strRes := `<DIV style='`// + sf.border.String() + `'></DIV>`
	return types.AHtml(strRes)
}

// GetStyle - -retur nstyle for DIV
func (sf *TFrame)GetStyle()types.AStyle{
	return types.AStyle(sf.border.String())
}
