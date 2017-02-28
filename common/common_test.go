package common

import (
	"testing"
)

func TestLog(t *testing.T) {
	InitLog()
	Log.Debug("test log")
	Log.Info("test log")
	Log.Error("test log")
}
