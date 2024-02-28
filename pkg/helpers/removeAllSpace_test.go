package helpers_test

import (
	"testing"

	"github.com/dmzsz/duozhuayu/pkg/helpers"
)

func TestRemoveAllSpace(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"Hello, World!", "Hello,World!"},
		{"    This   string   has   lots   of   spaces   ", "Thisstringhaslotsofspaces"},
		{"No spaces here", "Nospaceshere"},
		{"     ", ""},
		{"", ""},
	}

	for _, tc := range testCases {
		got := helpers.RemoveAllSpace(tc.input)
		if got != tc.want {
			t.Errorf("helpers.RemoveAllSpace(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}
