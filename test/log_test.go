package test

import (
	"ZTrunk_Server/logger"
	"testing"
)

func TestFileLogger(t *testing.T) {
	err := logger.InitLog("Test")
	if err != nil {
		panic(err)
		return
	}
	logger.Debug("Debug")
	logger.Warn("Warn")
	logger.Fatal("Fatal")
	logger.Error("Error")
	logger.Info("Info")
}
