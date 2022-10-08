package service

import (
	"gorm/authentication"
	
	"net/http"
)

//middleware to check if the request made is by an authorized user
func IsAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("token")
		if err != nil {
			w.Write([]byte("Login again.Token invalid"))
			return
		}

		okBool := authentication.VerifyJWT(cookie.Value)

		if !okBool {
			w.Write([]byte("Login again.Token invalid"))
			return
		}

		next.ServeHTTP(w, r)
	}
}


