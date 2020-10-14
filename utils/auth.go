package utils

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// AuthLayout is a date layout for auth
const AuthLayout = "20060102"

// Authenticate the request token
func Authenticate(r *http.Request) bool {
	username := os.Getenv("USERNAME")
	pass := os.Getenv("PASSWORD")
	if len(username) == 0 && len(pass) == 0 {
		fmt.Println("== Unaunthenticated ==")
		return true
	}
	reqToken := strings.TrimSpace(r.Header.Get("Authorization"))
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) == 2 {
		tokenStr := username + ":" + time.Now().Format(AuthLayout) + ":" + pass
		sysAuthToken := base64.StdEncoding.EncodeToString([]byte(tokenStr))
		return verify(sysAuthToken, splitToken[1])
	}
	return false
}

func verify(sysAuthToken, authToken string) bool {
	if authToken == sysAuthToken {
		return true
	}
	return false
}
