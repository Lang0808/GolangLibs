package config_test

import (
	"github.com/Lang0808/GolangLibs/config"

	"testing"
)

func TestGetProjectRoot(t *testing.T) {
	t.Log(config.GetProjectRoot())
}

func TestInit(t *testing.T) {
	config.Init("development")
}

func TestInitWithConfDir(t *testing.T) {
	config.InitWithConfDir("development", "../conf/")
}
