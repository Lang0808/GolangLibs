package config // import "github.com/Lang0808/GolangLibs/config"

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/spf13/viper"
)

func Init(mode string) error {
	if mode != "development" && mode != "production" {
		return errors.New("Mode is invalid")
	}
	confDir := GetConfDir()
	fmt.Println("Config directory: " + confDir)
	return ReadFilesInConfDir(confDir, mode)
}

func ReadFilesInConfDir(confDir string, mode string) error {
	confContents := ""
	err := filepath.Walk(confDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// only read files in confDir, directories are ignored
		if info.IsDir() {
			return nil
		}
		// fileName must match with mode and yaml
		matched, err := regexp.Match(mode+".*.yaml$", []byte(info.Name()))
		if matched {
			// read file content and add to confContents
			fmt.Println("Read file " + path)
			content, err := ioutil.ReadFile(path)
			if err == nil {
				confContents = confContents + string(content) + "\n"
			}
		}
		return nil
	})
	// use viper to read confContents
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer([]byte(confContents)))
	return err
}

func GetConfDir() string {
	projectRoot := GetProjectRoot()
	return projectRoot + "\\conf\\"
}

func GetProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
