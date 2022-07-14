package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtService struct {
	secretKey string
	issure    string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: "secret-key-mocada",
		issure:    "user-service",
	}
}

type Claim struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(name, email string) (string, error) {
	claim := &Claim{
		name,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), 
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("token invalido: %v", token)
		}
		return []byte(s.secretKey), nil
	})

	return err == nil
}
