package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateHashing(text string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return hashedPassword, err
}

func CheckPasswordHash(hashedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func GenerateToken(username string) (string, time.Time, error) {
	var jwtKey = []byte(os.Getenv("JWT_SECRET"))
	expirationTime := time.Now().Add(6 * 30 * 24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, expirationTime, err
}

func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	var jwtKey = []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, status.Errorf(codes.InvalidArgument, "invalid token")
	}

	if claims.ExpiresAt.After(time.Now()) {
		return nil, status.Errorf(codes.InvalidArgument, "expired token")
	}

	return claims, nil
}
