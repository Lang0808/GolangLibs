package mysql_test

import (
	"testing"

	"github.com/Lang0808/GolangLibs/mysql"
)

func TestFormatArrayMysql(t *testing.T) {

	a := []int{1, 2, 3, 4, 5}
	res := mysql.FormatArrayMysql[int](a)
	expected := "(1, 2, 3, 4, 5)"
	if res != expected {
		t.Fatalf("Expected %v Found %v\n", expected, res)
	}

	b := []int32{1, 2, 3, 4, 5}
	res = mysql.FormatArrayMysql[int32](b)
	expected = "(1, 2, 3, 4, 5)"
	if res != expected {
		t.Fatalf("Expected %v Found %v\n", expected, res)
	}

}
