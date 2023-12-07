package time

import (
	"strings"
	"time"
)

var loc *time.Location

var symbol2Value map[string]string

func init() {
	loc = time.FixedZone("Vietname", 7*3600)
	symbol2Value = make(map[string]string)
	symbol2Value["DD"] = "02"
	symbol2Value["MM"] = "01"
	symbol2Value["YYYY"] = "2006"
	symbol2Value["hh"] = "15"
	symbol2Value["mm"] = "04"
	symbol2Value["ss"] = "05"
}

// Parse time from string to int64 using layout, in time zone GMT+7
// Ex: Layout = "DD/MM/YYYY hh:mm:ss" and t = "07/12/2023 09:58:07", returns 1701917887000
func Parse(t string, layout string) (int64, error) {
	newLayout := layout
	for key, value := range symbol2Value {
		newLayout = strings.Replace(newLayout, key, value, -1)
	}
	a, err := time.ParseInLocation(newLayout, t, loc)
	if err != nil {
		return -1, err
	}
	return a.UnixMilli(), nil
}

// Convert time from int64 to string using layout, in time zone GMT+7
// Ex: Layout = "DD/MM/YYYY hh:mm:ss" and t = 1701917887000, returns "07/12/2023 09:58:07"
func Convert(t int64, layout string) string {
	newLayout := layout
	for key, value := range symbol2Value {
		newLayout = strings.Replace(newLayout, key, value, -1)
	}
	a := time.UnixMilli(t)
	// a := time.Unix(t, 7*3600)
	return a.Format(newLayout)
}
