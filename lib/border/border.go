package border

/*
	Пакет предоставляет потокобезопасный бордюр (универсальный элемент).
	Имеет параметры:
		- толщина
		- цвет
		- стиль
*/

import (
	"github.com/prospero78/goGui/lib/attr/line"
	"github.com/prospero78/goGui/lib/types"
)

// TBorder -- thread save operation with border
type TBorder struct {
	left   *line.TLine
	top    *line.TLine
	right  *line.TLine
	bottom *line.TLine
}

// NewBorder -- return new *TBorder
func NewBorder(color types.AColor, width types.AWidth, thickNess types.AThickness, lineStyle types.ALineStyle) (border *TBorder) {
	return &TBorder{
		left:   line.NewLine(width, thickNess, lineStyle, color),
		top:    line.NewLine(width, thickNess, lineStyle, color),
		right:  line.NewLine(width, thickNess, lineStyle, color),
		bottom: line.NewLine(width, thickNess, lineStyle, color),
	}
}

// Left -- return left line
func (sf *TBorder) Left() *line.TLine {
	return sf.left
}

// Top -- return top line
func (sf *TBorder) Top() *line.TLine {
	return sf.top
}

// Right -- return right line
func (sf *TBorder) Right() *line.TLine {
	return sf.right
}

// Bottom -- return bottom line
func (sf *TBorder) Bottom() *line.TLine {
	return sf.bottom
}

// SetThickness -- set thickness to all lines in border
func (sf *TBorder) SetThickness(thickNes types.AThickness) {
	sf.left.SetTihickness(thickNes)
	sf.top.SetTihickness(thickNes)
	sf.right.SetTihickness(thickNes)
	sf.bottom.SetTihickness(thickNes)
}
