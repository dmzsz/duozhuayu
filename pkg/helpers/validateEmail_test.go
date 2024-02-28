package helpers_test

import (
	"testing"

	"github.com/dmzsz/duozhuayu/pkg/helpers"
)

func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		email string
		want  bool
	}{
		{"test@example.com", true},
		{"in", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@gmail.com", false},
		{"invalid", false},
		{"invalid@", false},
		{"invalid@[127.0.0.1]", false},
		{"me@no-destination.pilinux.me", false},
		{"@missinglocalpart.com", false},
	}

	for _, tc := range testCases {
		got := helpers.ValidateEmail(tc.email)
		if got != tc.want {
			t.Errorf("helpers.ValidateEmail(%q) = %v, want %v", tc.email, got, tc.want)
		}
	}
}
