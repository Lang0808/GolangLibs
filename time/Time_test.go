package time_test

import (
	"testing"

	"github.com/Lang0808/GolangLibs/time"
)

func TestParse(t *testing.T) {
	list_test := []struct {
		desc     string
		t        string
		expected int64
	}{
		{"Test1", "07/12/2023 09:58:07", 1701917887000},
		{"Test2", "07/10/2022 09:00:00", 1665108000000},
		{"Test3", "abcdef", -1},
	}
	for _, test := range list_test {
		t.Run(test.desc, func(t *testing.T) {
			a, err := time.Parse(test.t, "DD/MM/YYYY hh:mm:ss")
			if err != nil && test.expected != -1 {
				t.Fatalf("%v; Unexpected Error: %v\n", test.desc, err)
			}
			if a != test.expected {
				t.Fatalf("%v; Expected %v found %v", test.desc, test.expected, a)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	list_test := []struct {
		desc     string
		expected string
		t        int64
	}{
		{"Test1", "07/12/2023 09:58:07", 1701917887000},
		{"Test2", "07/10/2022 09:00:00", 1665108000000},
		{"Test3", "01/01/1970 08:00:00", 0},
	}
	for _, test := range list_test {
		t.Run(test.desc, func(t *testing.T) {
			a := time.Convert(test.t, "DD/MM/YYYY hh:mm:ss")
			if a != test.expected {
				t.Fatalf("%v; Expected %v found %v", test.desc, test.expected, a)
			}
		})
	}
}
