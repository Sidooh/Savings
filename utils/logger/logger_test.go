package logger

import (
	"github.com/spf13/viper"
	"log/slog"
	"testing"
)

func TestLoggerInit(t *testing.T) {
	if !Log.Enabled(nil, slog.LevelInfo) {
		t.Errorf("Init() = info not enabled; want info level")
	}

	viper.Set("APP_ENV", "TEST")

	Init()

	if !Log.Enabled(nil, slog.LevelInfo) {
		t.Errorf("Init() = info not enabled; want info level")
	}
}
