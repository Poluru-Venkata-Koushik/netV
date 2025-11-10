package LogPackage

import (
	"log"
	"os"
	"sync"
)

const (
	Red            = "\033[31m"
	Green          = "\033[32m"
	Yellow         = "\033[33m"
	Blue           = "\033[34m"
	Reset          = "\033[0m"
	LogFileBasPath = "./Logs/"
)

var (
	Lock      sync.Mutex
	once      sync.Once
	ColorCode = map[string]string{
		"Error": Red,
		"Info": Blue,
		"CODE": Green,
		"Reset": Reset, // This is not needed. Adding it for sanity purposes.

	}
)

type Logger struct {
	// This is the Logger struct which will be used by all the go files to
	// log something to respective path
	Filename string
	PKG      string
}

func InitDIrectory() {
	// Has to be run only once. Hence we are using Sync.Once to ensure it
	// is only run once
	once.Do(
		func() {
			os.Mkdir(LogFileBasPath, 0755)
		})
}

/*
Errorlog method would be used to Log any errors
*/

func WritetoLog(filename string, toLog string, Status string, pkg string) {
	InitDIrectory()
	Lock.Lock()
	defer Lock.Unlock()
	file, _ := os.OpenFile(LogFileBasPath+filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	log.SetOutput(file)
	log.SetPrefix(" :: " + Status + " :: " + pkg + " ::")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	_ = log.Output(3, ColorCode[Status]+toLog+Reset)

}
func (L *Logger) ErrorLog(err string) {
	WritetoLog(L.Filename, err, "Error", L.PKG)
}

func (L *Logger) PrintLog(log string) {
	WritetoLog(L.Filename, log, "CODE", L.PKG)
}

func (L *Logger) InfoLog(log string) {
	WritetoLog(L.Filename, log, "Info", L.PKG)
}
