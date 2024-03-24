package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"online-judge/biz/model/entity"
	"time"
)

var secretKey = "cs_knowledge_base"

// UserClaims 自定义声明结构体
type UserClaims struct {
	UserID string `json:"user_id"`
	Exp    int64  `json:"exp"`
	jwt.RegisteredClaims
}

// CreateToken 创建访问令牌和刷新令牌
func CreateToken(userID string) (*entity.TokenPair, error) {
	accessClaims := generateClaims(userID, time.Second*60)
	accessToken, err := signToken(accessClaims)
	if err != nil {
		return nil, err
	}

	refreshClaims := generateClaims(userID, time.Hour*4)
	refreshToken, err := signToken(refreshClaims)
	if err != nil {
		return nil, err
	}

	return &entity.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// generateClaims 生成声明
func generateClaims(userID string, expiration time.Duration) *UserClaims {
	return &UserClaims{
		UserID: userID,
		Exp:    time.Now().Add(expiration).Unix(),
	}
}

// signToken 签名令牌
func signToken(claims *UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ParseTokenString 解析令牌
func ParseTokenString(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("token is not valid")
	}

	if claims.Exp < time.Now().Unix() {
		return nil, errors.New("token is expired")
	}

	if token.Valid {
		return claims, nil
	}

	return nil, errors.New("token is not valid")
}

// refreshATokenPair 续签访问令牌
func refreshATokenPair(refreshTokenString string) (*entity.TokenPair, error) {
	refreshClaims, err := ParseTokenString(refreshTokenString)
	if err != nil {
		return nil, err
	}

	// 生成新的访问令牌
	// generate a new access token
	newAccessClaims := generateClaims(refreshClaims.UserID, time.Hour*2)
	newAccessToken, err := signToken(newAccessClaims)
	if err != nil {
		return nil, err
	}

	// generate a new refresh token
	newRefreshClaims := generateClaims(refreshClaims.UserID, time.Hour*4)
	newRefreshToken, err := signToken(newRefreshClaims)
	if err != nil {
		return nil, err
	}

	return &entity.TokenPair{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
