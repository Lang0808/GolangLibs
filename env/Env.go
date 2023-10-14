package env // import "github.com/Lang0808/GolangLibs/env"

import (
	"os"

	"github.com/joho/godotenv"
)

var envs map[string]string

func InitEnv(envDir string) error {
	var err error
	envs, err = godotenv.Read(envDir)
	if err != nil {
		return err
	}
	return nil
}

// check if key exists in env file
// if not, check if key exists in env variable in server
func GetEnv(key string) string {
	val, ok := envs[key]
	if ok {
		return val
	}
	return os.Getenv(key)
}
