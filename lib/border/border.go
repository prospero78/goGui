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
	left   *line.TLineBorder
	top    *line.TLineBorder
	right  *line.TLineBorder
	bottom *line.TLineBorder
}

// NewBorder -- return new *TBorder
func NewBorder(color types.AColor, thickNess types.AThickness, lineStyle types.ALineStyle) (border *TBorder) {
	return &TBorder{
		left:   line.NewLineBorder(thickNess, lineStyle, color, "border-left"),
		top:    line.NewLineBorder(thickNess, lineStyle, color, "border-top"),
		right:  line.NewLineBorder(thickNess, lineStyle, color, "border-right"),
		bottom: line.NewLineBorder(thickNess, lineStyle, color, "boreder-bottom"),
	}
}

// Left -- return left line
func (sf *TBorder) Left() *line.TLineBorder {
	return sf.left
}

// Top -- return top line
func (sf *TBorder) Top() *line.TLineBorder {
	return sf.top
}

// Right -- return right line
func (sf *TBorder) Right() *line.TLineBorder {
	return sf.right
}

// Bottom -- return bottom line
func (sf *TBorder) Bottom() *line.TLineBorder {
	return sf.bottom
}

// SetThickness -- set thickness to all lines in border
func (sf *TBorder) SetThickness(thickNes types.AThickness) {
	sf.left.SetThickness(thickNes)
	sf.top.SetThickness(thickNes)
	sf.right.SetThickness(thickNes)
	sf.bottom.SetThickness(thickNes)
}

// String -- возвращает строковое представление бордюра (реализация интерфейса)
func (sf *TBorder) String() map[string]map[string]string {
	res := make(map[string]map[string]string)
	res["borderLeft"]=sf.left.GetStyle()
	res["borderRight"]=sf.right.GetStyle()
	res["borderTop"]=sf.top.GetStyle()
	res["borderBottom"]=sf.bottom.GetStyle()
	return res
}
