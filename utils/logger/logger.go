package logger

import (
	"Savings/utils"
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

var Log = slog.Default()

func Init() {
	Log = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	env := viper.GetString("APP_ENV")
	logger := viper.GetString("LOGGER")

	if env == "TEST" {
		return
	}

	if logger == "GCP" {
		//Format for GCP if needed
	} else {
		fileName := fmt.Sprintf("savings-%s.log", time.Now().Format("2006-01-02"))
		Log = slog.New(slog.NewJSONHandler(GetLogFile(fileName), nil))
	}

}

func GetLogFile(filename string) *os.File {
	pwd, err := os.Getwd()
	path := filepath.Join(pwd, "storage/logs/", filename)

	file := utils.GetFile(path)
	if err != nil || file == nil {
		file = utils.CreateFile(path)
	}

	return file
}
