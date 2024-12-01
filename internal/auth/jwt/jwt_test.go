package jwt

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {

	token, err := GenerateJWT("test-user")

	if err != nil {
		t.Errorf("expected error to be nil, but got %v", err)
	}

	if token == nil {
		t.Errorf("expected token to be present, but it was nil")
	} else {
		fmt.Printf("token: %s", *token)
	}
}
