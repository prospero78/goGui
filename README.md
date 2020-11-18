# goGui

-- ru --

Небольшая либа для создания графических интерфейсов (`GUI`) с помощью `Lorca` на `golang`.

Использует в качестве поверхности рисования окно `Chrome`

Должно работать на `Windos`, `Linux`, `MacOS`. Проверялось только на `Linux`

-- en --

A small lib for creating graphical interfaces (`GUI`) with `Lorca` on `golang`.

Uses the `Chrome` window as the drawing surface

Should work on `Windos`, `Linux`, `macOS`. Tested only on `Linux`

## Установка

Выполните следующую команду в консоли:

```bash
go get github.com/prospero78/goGui/...
```

Произойдёт скачивание всего пакета вместе с примерами (`./examples`) и графической библиотекой (`./lib`)

## Использование

Ниже представлен простейший пример:

```go
package main

/*
   Demonstrates a simple window.
*/

import(
   "github.com/prospero78/goGui/lib"
)

func main(){
   log:=lib.Log
   win, err:=lib.NewWin("Simple window")
   if err!=nil{
      log.Panicf("panic in create simple window\n\t%v", err)
   }
}
```

Другие примеры использования смотрите в папке `./examples`

## Makefile

```bash
make win.sample // Простой пример создания окна и манипуляции его свойствами.
```
