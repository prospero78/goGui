package linestyle

import (
	"sync"

	"github.com/prospero78/goGui/lib"
	"github.com/prospero78/goGui/lib/types"
)

/*
	Пакет предоставляет потокобезопасный тип для хранения стиля линии для отрисовки элемента.
*/

const (
	// None -- Линия не отображается и значение ее толщины обнуляется.
	None = types.ALineStyle("none")
	// Hidden -- Имеет тот же эффект, что и none за исключением применения border-bottom-style к ячейкам таблицы,
	//у которой значение свойства border-collapse установлено как collapse. В этом случае нижняя граница в ячейке не будет отображаться вообще.
	Hidden = types.ALineStyle("hidden")
	// Dotted -- Линия состоящая из набора точек.
	Dotted = types.ALineStyle("dotted")
	// Dashed -- пунктирная линия, состоящая из серии коротких отрезков.
	Dashed = types.ALineStyle("dashed")
	// Solid -- Сплошная линия.
	Solid = types.ALineStyle("solid")
	// Double -- Двойная линия.
	Double = types.ALineStyle("double")
	// Groove -- Создает эффект вдавленной линии на плоском элементе.
	Groove = types.ALineStyle("groove")
	// Ridge -- Создает эффект рельефной выпуклой линии на плоском элементе.
	Ridge = types.ALineStyle("ridge")
	// Inset -- Псевдотрехмерная вдавленная линия.
	Inset = types.ALineStyle("inset")
	// Outset -- Псевдотрехмерная выпуклая линия.
	Outset = types.ALineStyle("outset")
	// Inherit -- Наследует значение родителя.
	Inherit = types.ALineStyle("inherit")
)

// TLineStyle -- thread save opration with line style
type TLineStyle struct {
	val   types.ALineStyle
	block sync.RWMutex
}

// NewLineStyle -- return new *TLineStyle
func NewLineStyle(style types.ALineStyle) (ls *TLineStyle) {
	ls = &TLineStyle{}
	ls.SetLineStyle(style)
	return ls
}

// SetLineStyle -- set line style
func (sf *TLineStyle) SetLineStyle(style types.ALineStyle) {
	defer sf.block.Unlock()
	sf.block.Lock()
	switch style {
	case None, Hidden, Dotted, Dashed, Solid, Double, Groove, Ridge, Inset, Outset, Inherit:
	default:
		lib.Log.Errorf("TLineStyle.SetLineStyle(): unknown line style(%v)\n", style)
		style = Inherit
	}
	sf.val = style
}

// GetLineStyle -- return saved line style
func (sf *TLineStyle) GetLineStyle() types.ALineStyle {
	defer sf.block.RUnlock()
	sf.block.RLock()
	return sf.val
}
