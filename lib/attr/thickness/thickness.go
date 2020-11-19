package thickness

/*
	Пакет предоставляет потокобезопасную толщину для абстрактного элемента.
*/

import (
	"fmt"
	"sync"

	"github.com/prospero78/goGui/lib/types"
)

// TThickness -- thread save operation with thickness
type TThickness struct {
	val   types.AThickness
	block sync.RWMutex
}

// NewThickness -- return new *TThickness
func NewThickness(val types.AThickness) (width *TThickness) {
	return &TThickness{
		val: val,
	}
}

// SetTihickness -- set thickness
func (sf *TThickness) SetTihickness(thickness types.AThickness) {
	defer sf.block.Unlock()
	sf.block.Lock()
	sf.val = thickness
}

// GetThickness -- return saved thickness
func (sf *TThickness) GetThickness() types.AThickness {
	defer sf.block.RUnlock()
	sf.block.RLock()
	return sf.val
}

// String -- возвращает толщину линии (реализация интерфейса)
//Поскольку мы не знаем .к чему относится толщина -- атрибут не преедаём.
func (sf *TThickness) String() string {
	defer sf.block.RUnlock()
	sf.block.RLock()
	return fmt.Sprint(sf.val)
}
