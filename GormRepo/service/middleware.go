package service

import (
	"gorm/jwtauthentication"
	"net/http"
)

func IsAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("token")
		if err != nil {
			w.Write([]byte("Login again.Token invalid"))
			return
		}

		okBool := jwtauthentication.VerifyJWT(cookie.Value)

		if !okBool {
			w.Write([]byte("Login again.Token invalid"))
			return
		}

		next.ServeHTTP(w, r)
	}
}


