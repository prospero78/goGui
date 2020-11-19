package line

import (
	"github.com/prospero78/goGui/lib/attr/color"
	"github.com/prospero78/goGui/lib/attr/linestyle"
	"github.com/prospero78/goGui/lib/attr/thickness"
	"github.com/prospero78/goGui/lib/types"
)

/*
	Пакет предоставляет тип линии с толщиной, цветом и стилем.
*/

// TLineBorder -- thread save operation with line
type TLineBorder struct {
	thickness *thickness.TThickness // Толщина линии
	lineStyle *linestyle.TLineStyle // Стиль линии
	color     *color.TColor         // Цвет линии
	typeLine  types.ALineType       // Тип линии (для толщины)
}

// NewLineBorder -- return new &TLine
func NewLineBorder(
	thickNess types.AThickness,
	lineStyle types.ALineStyle,
	colorLine types.AColor,
	typeLine types.ALineType) (line *TLineBorder) {
	line = &TLineBorder{
		thickness: thickness.NewThickness(thickNess),
		lineStyle: linestyle.NewLineStyle(lineStyle),
		color:     color.NewColor(colorLine),
		typeLine:  typeLine,
	}
	return line
}

// GetStyle -- возвращает словарь стиля
func (sf *TLineBorder) GetStyle() map[string]string {
	thickness := sf.thickness.String()
	lineStyle := sf.lineStyle.String()
	color := sf.color.String()
	res := make(map[string]string)
	res["Color"] = color
	res["Style"] = lineStyle
	res["Width"] = thickness
	return res
}

// SetThickness -- устанавливает толщину линии
func (sf *TLineBorder) SetThickness(thickNess types.AThickness) {
	sf.thickness.SetTihickness(thickNess)
}
