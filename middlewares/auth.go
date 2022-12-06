package middlewares

import (
	"airline/utils"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func Auth(handler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")

		if err != nil {
			result, _ := json.Marshal(map[string]string{
				"message": "unauthorized",
			})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(result)
			return
		}

		token, claims, err := utils.ParseToken(cookie)
		if err != nil {
			result, _ := json.Marshal(map[string]string{
				"message": "unauthorized",
			})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(result)
			return
		}
		if !token.Valid {
			result, _ := json.Marshal(map[string]string{
				"message": "token is not valid",
			})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(result)
			return
		}

		idClaim, _ := strconv.ParseInt(claims["iss"].(string), 10, 64)
		idCtx := context.WithValue(r.Context(), "idFromToken", idClaim)

		handler.ServeHTTP(w, r.WithContext(idCtx))
	}

	return http.HandlerFunc(fn)
}
