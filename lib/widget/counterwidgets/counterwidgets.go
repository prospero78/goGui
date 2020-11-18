package counterwidgets

/*
	Потокобезопасный глобальный счётчик виджетов, один на всех. Только увеличивается.
*/

import (
	"sync"

	"github.com/prospero78/goGui/lib/types"
)

// TCounterWidgets -- потокобезопасные операции с глобальным счётчиком виджетов
type TCounterWidgets struct {
	val   types.AWidgetID
	block sync.RWMutex
}

var (
	//CounterWidgets -- глобальный потокобезопасный счётчик виджетов
	CounterWidgets *TCounterWidgets
)

// Inc -- add 1 for global counter widgets
func (sf *TCounterWidgets) Inc() {
	defer sf.block.Unlock()
	sf.block.Lock()
	sf.val++
}

// Get -- return saved counter widgets
func (sf *TCounterWidgets) Get() types.AWidgetID {
	defer sf.block.RUnlock()
	sf.block.RLock()
	return sf.val
}

func init() {
	CounterWidgets = &TCounterWidgets{}
}
