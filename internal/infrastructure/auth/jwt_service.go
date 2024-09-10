// internal/infrastructure/auth/jwt_service.go
package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService(secretKey, issuer string) JWTService {
	return &jwtService{
		secretKey: secretKey,
		issuer:    issuer,
	}
}

func (s *jwtService) GenerateToken(userID string) (string, error) {
	claims := &jwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    s.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
}
