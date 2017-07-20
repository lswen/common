package log

import (
	"testing"
)

func Test_Log(t *testing.T) {
	Init("DEVELOPMENT")
	Debug("debug")
	Info("info")
	Warn("warn")
}
