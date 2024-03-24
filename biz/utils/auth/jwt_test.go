package auth

import "testing"

func TestCreateToken(t *testing.T) {
	token, err := CreateToken("123456")
	if err != nil {
		t.Errorf("CreateToken() error = %v", err)
		return
	}

	t.Log("token:", token)
}

func TestParseTokenString(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDU2IiwiZXhwIjoxNzExMjg3NDkyfQ.eXAAx0Tg1PIcngl7B9Ym5g-cEnsmQY0--d3BTVIJiEI"
	claims, err := ParseTokenString(tokenString)
	if err != nil {
		t.Errorf("ParseTokenString() error = %v", err)
		return
	}

	t.Log("claims:", claims.UserID)
}
