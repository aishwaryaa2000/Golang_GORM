package authentication

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("your-256-bit-secret")

func GenerateJWT(id string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["authorized"] = true
	claims["id"] = id

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		fmt.Println("Error is : ", err)
		return "", err
	}

	fmt.Println("Token inside jwt auth : ", tokenString)
	return tokenString, nil

}

//verifyJWT function returns true if token is valid else returns false.
func VerifyJWT(tokenString string) bool {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing token")
		}

		return sampleSecretKey, nil
	})

	if err != nil {
		fmt.Println("Invalid token ", err)
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}

	return false

}
