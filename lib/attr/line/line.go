package line

import (
	"github.com/prospero78/goGui/lib/attr/color"
	"github.com/prospero78/goGui/lib/attr/linestyle"
	"github.com/prospero78/goGui/lib/attr/thickness"
	"github.com/prospero78/goGui/lib/attr/width"
	"github.com/prospero78/goGui/lib/types"
)

/*
	Пакет предоставляет тип линии с толщиной, цветом и стилем.
*/

// TLine -- thread save operation with line
type TLine struct {
	*width.TWidth         // Длина линии
	*thickness.TThickness // Толщина линии
	*linestyle.TLineStyle // Стиль линии
	*color.TColor         // Цвет линии
}

// NewLine -- return new &TLine
func NewLine(
	widthLine types.AWidth,
	thickNess types.AThickness,
	lineStyle types.ALineStyle,
	colorLine types.AColor) (line *TLine) {
	line = &TLine{
		TWidth:     width.NewWidth(widthLine),
		TThickness: thickness.NewThickness(thickNess),
		TLineStyle: linestyle.NewLineStyle(lineStyle),
		TColor:     color.NewColor(colorLine),
	}
	return line
}
