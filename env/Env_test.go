package env_test

import (
	"os"
	"testing"

	"github.com/Lang0808/GolangLibs/env"
)

func TestGetEnv(t *testing.T) {
	err := env.InitEnv("../.env")
	os.Setenv("TEST3", "server testtest3")
	os.Setenv("TEST4", "server testtest4")
	if err != nil {
		// init env fail
		t.Fatal(err)
	}
	list_test := []struct {
		desc     string
		key      string
		expected string
	}{
		{"in file env 1", "TEST", "testtest"},
		{"in file env 2", "TEST2", "testtest2"},
		{"in file env and in server variable", "TEST3", "testtest3"},
		{"not in file env and in server variable", "TEST4", "server testtest4"},
		{"not in file env and not in server variable", "TEST5", ""},
	}

	for _, test := range list_test {
		t.Run(test.desc, func(t *testing.T) {
			val := env.GetEnv(test.key)
			if val != test.expected {
				t.Fatalf("%v; expected = %v, found = %v;\n", test.desc, test.expected, val)
			}
		})
	}
}
