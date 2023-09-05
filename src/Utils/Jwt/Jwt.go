package Jwt

import (
	"os"
	"time"

	// "github.com/gofiber/storage/mysql"

	"github.com/golang-jwt/jwt/v4"
)


func GenerateJwtToken() string {
	jwtSecret := os.Getenv("JWT_SECRET")
	token  := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Topan"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		panic(err)
	}
	
	// return tokenString
	return tokenString
}

func VerifyJwtToken(token string) (jwt.MapClaims, bool) {
	 // Parse the token
	 jwtSecret := os.Getenv("JWT_SECRET")
	 t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		 return []byte(jwtSecret), nil
	 })
	 if err != nil {
		 return nil, false
	 }
	 claims, ok := t.Claims.(jwt.MapClaims)
	 if !ok || !t.Valid {
		 return nil, false
	 }
	 return claims, true
}