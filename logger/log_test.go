package logger

import "testing"

func TestConsoleLog(t *testing.T) {
	log := CreateConsoleLog(DEBUG)
	log.Debug("Debug\n")
	log.Trace("Trace\n")
	log.Info("Info\n")
	log.Error("Error\n")
	log.Warn("Warn\n")
	log.Fatal("Fatal")
}
