package log

/*
	The simplest logging library provides several features that the standard
	library logging implementation does not have.
*/

import (
	"fmt"
	"time"
)

const (
	// WARN -- level for warning
	WARN = iota
	// INFO -- level for information
	INFO
	// DEBUG -- set level debug
	DEBUG
)

// TLog -- operations with message logging
type TLog struct {
	prefix string
	level  int
}

// NewLog -- returns new *TLog
func NewLog() *TLog {
	return &TLog{}
}

// SetPrefix -- set prefix for output
func (sf *TLog) SetPrefix(prefix string) {
	sf.prefix = prefix
}

// SetLevel -- set level for output
func (sf *TLog) SetLevel(level int) {
	if level < 0 {
		sf.Errorf("level must by %v..%v", DEBUG, WARN)
		level = 0
	}
	sf.level = level
}

//Infof -- out info message
func (sf *TLog) Infof(msg string, arg ...interface{}) {
	if sf.level >= INFO {
		if arg != nil {
			fmt.Printf("INFO %v %v.%v %+v\n", sf.getTime(), sf.prefix, msg, arg)
			return
		}
		fmt.Printf("INFO %v %v.%v\n", sf.getTime(), sf.prefix, msg)
	}
}

//Errorf -- out error message
func (sf *TLog) Errorf(msg string, arg ...interface{}) {
	if arg != nil {
		fmt.Printf("ERRO %v %v%v %+v\n", sf.getTime(), sf.prefix, msg, arg)
		return
	}
	fmt.Printf("ERRO %v %v.%v\n", sf.getTime(), sf.prefix, msg)
}

//Panicf -- out info message
func (sf *TLog) Panicf(msg string, arg ...interface{}) {
	if arg != nil {
		fmt.Printf("PANI %v %v.%v %+v\n", sf.getTime(), sf.prefix, msg, arg)
		panic(nil)
	}
	fmt.Printf("PANI %v %v.%v\n", sf.getTime(), sf.prefix, msg)
	panic(nil)
}

// Return current time to out log
func (sf *TLog) getTime() (res string) {
	res = time.Now().UTC().Format("2006-01-02 15:04:05.000")
	return res
}
