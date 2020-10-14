package utils_test

import (
	"encoding/base64"
	"net/http"
	"os"
	"testing"
	"time"

	"example.com/emailreports/utils"
)

var authTestcaese = []struct {
	name, username, password     string
	wrongPass, op, withoutbearer bool
}{
	{
		name:     "No authentication",
		username: "",
		password: "",
		op:       true,
	},
	{
		name:     "Authenticated",
		username: "test",
		password: "test",
		op:       true,
	},
	{
		name:      "Unauthenticated",
		username:  "test",
		password:  "test",
		wrongPass: true,
		op:        false,
	},
	{
		name:          "Unauthenticated-Improper-Token",
		username:      "test",
		password:      "test",
		withoutbearer: true,
		op:            false,
	},
}

func TestAuthenticate(t *testing.T) {
	t.Parallel()
	genToken := func(username, pass string) string {
		tokenStr := username + ":" + time.Now().Format(utils.AuthLayout) + ":" + pass
		sysAuthToken := base64.StdEncoding.EncodeToString([]byte(tokenStr))
		return sysAuthToken
	}

	for _, tc := range authTestcaese {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("USERNAME", tc.username)
			os.Setenv("PASSWORD", tc.password)
			if tc.wrongPass {
				tc.password = "newPassword"
			}
			var r = http.Request{}
			r.Header = make(map[string][]string)
			bearer := "Bearer"
			if tc.withoutbearer {
				bearer = ""
			}
			r.Header.Set("Authorization", bearer+" "+genToken(tc.username, tc.password))
			got := utils.Authenticate(&r)
			want := tc.op
			if got != want {
				t.Error()
			}
			os.Unsetenv("USERNAME")
			os.Unsetenv("PASSWORD")
		})
	}
}
