package log_test

import (
	"testing"

	"github.com/Lang0808/GolangLibs/config"
	"github.com/Lang0808/GolangLibs/log"
)

func TestInitLog(t *testing.T) {
	config.InitWithConfDir("development", "../conf/")
	err := log.InitLog()
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}
	log.Error("Test log error")
	log.Error("Test log error 2")
	log.Info("Test log info")
}
