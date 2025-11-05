package main

import (
	Logging "netVServer/LogPackage"
)

const (
	LogFile = "MongoClient.log"
	PKG     = "Database"
)

var (
	log = Logging.Logger{Filename: LogFile, PKG: PKG}
)

func main() {
	log.ErrorLog("Game start!!")
}
