package auth

import (
	"testing"
)

func TestJWT(t *testing.T) {
	token, err := CreateToken("123456")
	if err != nil {
		t.Errorf("CreateToken() error = %v", err)
		return
	}

	t.Log("token:", token)

	claims, err := ParseTokenString(token.AccessToken)
	if err != nil {
		t.Errorf("ParseTokenString() error = %v", err)
		return
	}

	t.Log("claims:", claims.UserID)

	refresh, err := RefreshATokenPair(token.RefreshToken)
	if err != nil {
		t.Errorf("RefreshATokenPair() error = %v", err)
		return
	}

	t.Log("refresh:", refresh.AccessToken)
}
