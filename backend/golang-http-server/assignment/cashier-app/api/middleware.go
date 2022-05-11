package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const (
	roleAdmin    = "admin"
	roleEmployee = "employee"
)

func (api *API) AllowOrigin(w http.ResponseWriter, req *http.Request) {
	// localhost:9000 origin mendapat ijin akses
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9000")
	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	// allow cookie
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
}

func (api *API) AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.AllowOrigin(w, r)
		encoder := json.NewEncoder(w)
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//       2. return unauthorized ketika token kosong
		//       3. return bad request ketika field token tidak ada
		tokenCookie, err := r.Cookie(jwtCookieKey)
		if err != nil {
			if err == http.ErrNoCookie {
				// return unauthorized ketika token kosong
				w.WriteHeader(http.StatusUnauthorized)
				encoder.Encode(AuthErrorResponse{"token is empty"})
				return
			}
			// return bad request ketika field token tidak ada
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(AuthErrorResponse{"token field not exists"})
			return
		}

		// Task: Ambil value dari cookie token
		tokenValue := tokenCookie.Value

		// Task: Deklarasi variable claim
		claims := &Claims{}

		// Task: 1. parse JWT token ke dalam claim
		//       2. return unauthorized ketika signature invalid
		//       3. return bad request ketika field token tidak ada
		//       4. return unauthorized ketika token sudah tidak valid (biasanya karna token expired)
		token, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				encoder.Encode(AuthErrorResponse{"invalid signature"})
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(AuthErrorResponse{"token field not exists"})
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(AuthErrorResponse{"invalid/expired token"})
			return
		}

		// Task: Validasi
		ctx := context.WithValue(r.Context(), claimsCtxKey, claims)
		ctx = context.WithValue(ctx, usernameCtxKey, claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		ctx := r.Context()
		claims := ctx.Value(claimsCtxKey).(*Claims)
		if claims.Role != roleAdmin {
			w.WriteHeader(http.StatusForbidden)
			encoder.Encode(AuthErrorResponse{"need admin role"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (api *API) GET(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.AllowOrigin(w, r)
		encoder := json.NewEncoder(w)
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			encoder.Encode(AuthErrorResponse{Error: "Need GET Method!"})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (api *API) POST(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.AllowOrigin(w, r)
		encoder := json.NewEncoder(w)
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			encoder.Encode(AuthErrorResponse{Error: "Need POST Method!"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
