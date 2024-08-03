/*
Could do something fancier with the logging, for example injecting the
logger to use rather than it being set, but YAGNI in this instance.
Simple + done rather than looking for abstractions/excuses to use patterns
when I don't need it.
*/

package server

import (
	"log"
	"os"
	"sync"
)

var l *log.Logger
var lock = &sync.Mutex{}

func getLogger() *log.Logger {
	if l == nil {
		lock.Lock()
		defer lock.Unlock()
		if l == nil {
			l = log.New(os.Stdout, "", log.LstdFlags)
		}
	}

	return l
}

func LogError(err error) {
	l = getLogger()
	l.SetPrefix("Server Error: ")
	l.Println(err)
}

func LogInfo(msg string) {
	l = getLogger()
	l.SetPrefix("Server Info: ")
	l.Println(msg)
}
