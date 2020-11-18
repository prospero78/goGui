package widgetid

import "github.com/prospero78/goGui/lib/types"

/*
	Пакет предоставляет потокобезопасный индивидуальный ID для каждого TWidget.
*/

// TWidgetID -- thread save operation for ID widgets
type TWidgetID struct {
	val types.AWidgetID
}

// NewWidgetID -- return new *TWidgetId
func NewWidgetID(id types.AWidgetID) (wid *TWidgetID) {
	return &TWidgetID{
		val: id,
	}
}

// GetWidgetID -- return saved widget ID
func (sf *TWidgetID) GetWidgetID() types.AWidgetID {
	return sf.val
}
