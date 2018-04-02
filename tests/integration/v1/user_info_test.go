package tests

import (
	"testing"
)

func TestUserInfo(t *testing.T) {
	_, err := client.User.Info()

	if err != nil {
		t.Fatalf("User.Info() returned error: %v", err)
	}
}
