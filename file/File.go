package file // import "github.com/Lang0808/GolangLibs/file"

import (
	"errors"
	"os"
)

func OpenOrCreateFile(file_path string) (*os.File, error) {

	path_seperator := '/'
	seperate_index := -1

	for i := len(file_path) - 1; i >= 0; i -= 1 {
		if file_path[i] == byte(path_seperator) {
			seperate_index = i
			break
		}
	}

	if seperate_index > 0 {
		only_folder := file_path[:seperate_index]
		err := os.MkdirAll(only_folder, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func IsExist(file_path string) bool {
	if _, err := os.Stat(file_path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
