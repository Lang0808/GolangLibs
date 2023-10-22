package jwt_test

import (
	"testing"

	"github.com/Lang0808/GolangLibs/config"
	"github.com/Lang0808/GolangLibs/jwt"
)

func TestCreateJWTToke(t *testing.T) {
	err := config.InitWithConfDir("development", "../conf/")
	if err != nil {
		t.Fatalf("Init config fail. Error=%v\n", err)
	}
	list_test := []struct {
		desc   string
		UserId string
	}{
		{"Test1", "sasounxiuqwwqbxiywq"},
		{"Test2", "qwkuxqywqxutvudeq"},
	}
	for _, test := range list_test {
		t.Run(test.desc, func(t *testing.T) {
			tok, err := jwt.CreateJWTToken(test.UserId)
			if err != nil {
				t.Fatalf("%v; CreateJWTToken; Error=%v\n", test.desc, err)
			}
			uid, err := jwt.GetUserIdInJWTToken(tok)
			if err != nil {
				t.Fatalf("%v; GetUserIdInJWTToken; Error=%v\n", test.desc, err)
			}
			if uid != test.UserId {
				t.Fatalf("%v; expected = %v, found = %v;\n", test.desc, test.UserId, uid)
			}
		})
	}
}
