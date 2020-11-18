package size

import (
	"sync"

	"github.com/prospero78/goGui/lib"
)

/*
	The package provides a type for operations on the size of elements
*/

// ASizeX -- ad hoc type for horizontal dimension type
type ASizeX int

// ASizeY -- ad hoc type for vertical dimension type
type ASizeY int

// TSize -- operation with size graphical elements
type TSize struct {
	x       ASizeX // size X (horizontal)
	y       ASizeY //size Y (vertical)
	isFixed bool   // Fized size
	block   sync.RWMutex
}

// NewSize - -return new *TSize
func NewSize(x ASizeX, y ASizeY) (size *TSize) {
	{ // Precondition
		if x < 0 {
			lib.Log.Panicf("NewSize(): x(%v)<0\n", x)
		}
		if y < 0 {
			lib.Log.Panicf("NewSize(): y(%v)<0\n", y)
		}
	}
	size = &TSize{
		x: x,
		y: y,
	}
	return size
}

// GetX -- return saved coord X
func (sf *TSize) GetX() ASizeX {
	defer sf.block.RUnlock()
	sf.block.RLock()
	return sf.x
}

// GetY -- return saved coord Y
func (sf *TSize) GetY() ASizeY {
	defer sf.block.RUnlock()
	sf.block.RLock()
	return sf.y
}

// SetFixed -- lock size for elements
func (sf *TSize) SetFixed() {
	defer sf.block.Unlock()
	sf.block.Lock()
	sf.isFixed = true
}

// ResetFixed -- unlock size for elements
func (sf *TSize) ResetFixed() {
	defer sf.block.Unlock()
	sf.block.Lock()
	sf.isFixed = false
}

// IsFixed - -return isFixed for this elements
func (sf *TSize) IsFixed() bool{
	defer sf.block.RUnlock()
	sf.block.RLock()
	return sf.isFixed
}
