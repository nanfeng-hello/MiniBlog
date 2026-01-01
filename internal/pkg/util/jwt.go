package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nanfeng/mini-blog/internal/config"
)

func GenerateToken(id string) (string, error) {

	// 1.获取配置信息
	cfg := config.Cfg.JwtConfig

	// 1.构建 jwt.Register
	claims := jwt.RegisteredClaims{
		Issuer:    cfg.Iss,
		Subject:   id,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.Exp) * time.Minute)),
	}

	// 2.进行base编码
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3.使用密钥进行加密
	return token.SignedString([]byte(cfg.Secret))
}

// ParseToken
func ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	// 1.使用jwt中的方法进行解秘
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(config.Cfg.JwtConfig.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	// 2.判断 tokenString 是否有效
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
