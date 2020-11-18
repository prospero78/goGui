package lib

/*
	The library provides global objects. Customizes components as needed.
*/

import(
	"github.com/prospero78/goGui/lib/log"
)

var (
	// Log -- global log object
	Log *log.TLog
)

// Initialization package at starttup
func init(){
	Log=log.NewLog()
	Log.SetLevel(log.DEBUG)
}
