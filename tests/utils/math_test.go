package utils_test

import (
	"testing"

	"example.com/emailreports/utils"
)

func TestRoundFloat(t *testing.T) {
	t.Parallel()
	ip := 1.2345
	got := utils.RoundFloat(ip)
	want := 1.23
	if got != want {
		// t.Errorf("\n\n ===> TestRoundFloat failed <===\n\n got  ===> %v \n\n want ===> %v \n\n", got, want)
		t.Error()
	}
}
