package authentication

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

func GetIdFromCookieClaims(r *http.Request) uuid.UUID {
	cookie, err := r.Cookie("token")
	if err != nil {
		log.Fatal(err)
	}
	tokenString := cookie.Value

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing token")
		}

		return sampleSecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(string)
		uuidId, _ := uuid.FromString(id)
		return uuidId

	}

	return uuid.Nil
}

func SetCookieValue(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Minute * 5),
	})
}


func DeleteCookieValue(w http.ResponseWriter){
	http.SetCookie(w,&http.Cookie{
		Name : "token",
		Value : " ",
		MaxAge:   -1,
		// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	})
}