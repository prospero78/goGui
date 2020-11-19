package color

import (
	"sync"

	"github.com/prospero78/goGui/lib/types"
)

/*
	Пакет предоставляет потокобезопасный тип цвета для абстрактного элемента.
*/

// TColor -- thread save operation with color
type TColor struct {
	val   types.AColor
	block sync.RWMutex
}

// NewColor -- return new *TColor
func NewColor(clr types.AColor) (color *TColor) {
	return &TColor{
		val: clr,
	}
}

// GetColor -- return save color
func (sf *TColor) GetColor() types.AColor {
	defer sf.block.RUnlock()
	sf.block.RLock()
	return sf.val
}

// SetColor -- set save color
func (sf *TColor) SetColor(color types.AColor) {
	defer sf.block.Unlock()
	sf.block.Lock()
	sf.val = color
}

// String -- return string represent color (realise interface)
func (sf *TColor) String() string {
	defer sf.block.RUnlock()
	sf.block.RLock()
	return string(sf.val)
}
