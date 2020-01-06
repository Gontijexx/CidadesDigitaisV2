package middlewares

import (
	"errors"
	"net/http"

	"CidadesDigitaisV2/api/responses"
	"CidadesDigitaisV2/api/auth"

)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, log.Printf("[FATAL] Unauthorized"))
			return
		}
		next(w, r)
	}
}