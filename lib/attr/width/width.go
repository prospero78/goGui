package width

/*
	Пакет предоставляет потокобезопасную длину для абстрактного элемента.
*/

import (
	"sync"

	"github.com/prospero78/goGui/lib/types"
)

// TWidth -- thread save operation with width
type TWidth struct {
	val   types.AWidth
	block sync.RWMutex
}

// NewWidth -- return new *TWidth
func NewWidth(val types.AWidth) (width *TWidth) {
	return &TWidth{
		val: val,
	}
}

// SetWidth -- set width
func (sf *TWidth) SetWidth(width types.AWidth) {
	defer sf.block.Unlock()
	sf.block.Lock()
	sf.val = width
}

// GetWidth -- return saved width
func (sf *TWidth) GetWidth() types.AWidth {
	defer sf.block.RUnlock()
	sf.block.RLock()
	return sf.val
}
