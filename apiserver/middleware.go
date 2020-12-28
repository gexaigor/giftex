package apiserver

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gexaigor/MyRestAPI/model"
)

// jwtMiddleware ...
func (s *Server) jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/account", "/auth"}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath && r.Method == "POST" {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			s.error(w, r, http.StatusForbidden, errorMissingToken)
			return
		}

		tk := &model.Token{}

		token, err := jwt.ParseWithClaims(tokenHeader, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.secretKey), nil
		})

		if err != nil {
			s.error(w, r, http.StatusForbidden, errorMalfordedToken)
			return
		}

		if !token.Valid {
			s.error(w, r, http.StatusForbidden, errorTokenIsNotValid)
			return
		}

		user, err := s.store.User().FindByID(tk.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				s.error(w, r, http.StatusNotFound, errorNotFound)
				return
			}
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		ctx := context.WithValue(r.Context(), ctxKeyUser, user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// userMiddleware ...
func (s *Server) userMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(ctxKeyUser).(*model.User)
		if !ok {
			s.error(w, r, http.StatusInternalServerError, errorConvertation)
			return
		}

		if user.Role != model.USER {
			s.error(w, r, http.StatusForbidden, errorForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// companyMiddleware ...
func (s *Server) companyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(ctxKeyUser).(*model.User)
		if !ok {
			s.error(w, r, http.StatusInternalServerError, errorConvertation)
			return
		}

		if user.Role != model.COMPANY {
			s.error(w, r, http.StatusForbidden, errorForbidden)
			return
		}

		company, err := s.store.Company().FindByUser(user)
		if err != nil {
			if err == sql.ErrNoRows {
				s.error(w, r, http.StatusNotFound, errorNotFound)
				return
			}
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		ctx := context.WithValue(r.Context(), ctxKeyUser, company)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
