package logger

import (
	"testing"
)

/*
func TestConsoleLog(t *testing.T) {
	log := CreateConsoleLog(DEBUG)
	log.Debug("Debug\n")
	log.Trace("Trace\n")
	log.Info("Info\n")
	log.Error("Error\n")
	log.Warn("Warn\n")
	log.Fatal("Fatal")
}
*/

func TestFileLogger(t *testing.T) {
	err := InitLog("Beyond")
	if err != nil {
		panic(err)
		return
	}
	log.Debug("Debug")
	log.Warn("Warn")
	log.Fatal("Fatal")
	log.Error("Error")
	log.Info("Info")
}
