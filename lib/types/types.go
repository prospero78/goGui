package types

/*
	-ru-
	Пакет предоставляет интерфейсные типы для развязки различных частей библиотеки

	-en-
	The package provides interface types to decouple various parts of the library
*/

// AWidgetID -- ad hoc types for widgets ID
type AWidgetID uint64

// AWidth -- ad hoc types for width elements
type AWidth uint64

// AThickness -- ad hoc types for thickness elements
type AThickness uint64

// AColor -- ad hoc types for color elements
type AColor string

// ALineStyle -- ad hoc types for line style in elements
type ALineStyle string

// AHtml -- ad hoc types for HTML-represent elements
type AHtml string

// IWidget -- универсальный интерфейс для всех виджетов.
type IWidget interface {
	// GetWidgetID -- возвращает ID виджета
	GetWidgetID() AWidgetID
	// GetHTML -- возвращает HTML-представление виджета
	GetHTML() AHtml
}

// IParent -- универсальный интерфейс предка
type IParent interface {
	// AddWidget -- добавляет в себя виджет
	AddWidget(IWidget)
}
