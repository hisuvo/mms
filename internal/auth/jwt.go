package auth

import (
	"fmt"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

//jwt secret key is loaded from env file in production
//for testing purpose it is hardcoded to "supersecretkey" in cmd/main.go
//secret key should be long enough (at least 256 bits) for HS256 algorithm
//jwt duration is 24 hours

type JWT struct {
	SecretKey []byte
	Duration  time.Duration
}

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"user_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type JwtService interface {
	GenerateToken(usreId string, email string, username string, role string) (string, error)
	ValidateToken(tokenstring string) (*Claims, error)
}


func NewJWT(secretKey []byte, duration time.Duration) JwtService {
	return &JWT{
		SecretKey: secretKey,
		Duration:  duration,
	}
}

func (j *JWT) GenerateToken(usreId string, email string, username string, role string) (string, error) {
	claims := &Claims{
		UserID: usreId,
		Email: email,
		Username: username,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.Duration)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "mms",
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SecretKey)
}


func (j *JWT) ValidateToken(tokenstring string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.SecretKey, nil
	})

	if err == nil && token.Valid {
		return token.Claims.(*Claims), nil
	}
	return nil, err
}