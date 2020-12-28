package apiserver

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gexaigor/MyRestAPI/model"
	"github.com/gexaigor/MyRestAPI/service"
	"github.com/lib/pq"
)

//------------ACCOUNT------------
// handleAccountCreate ...
func (s *Server) handleAccountCreate() http.HandlerFunc {
	type request struct {
		Login     string `json:"login"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		IsCompany bool   `json:"isCompany"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		role := model.USER
		if req.IsCompany {
			role = model.COMPANY
		}

		u := &model.User{
			Login:    req.Login,
			Email:    req.Email,
			Password: req.Password,
			Role:     role,
		}

		err := s.store.User().Save(u)
		if err != nil {
			if pgerr, ok := err.(*pq.Error); ok {
				if pgerr.Code == pgUniqueViolation {
					s.error(w, r, http.StatusBadRequest, errorUniqueViolation)
					return
				}
			}
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if u.Role == model.COMPANY {
			company := &model.Company{
				User: u,
			}
			err := s.store.Company().Save(company)
			if err != nil {
				if pgerr, ok := err.(*pq.Error); ok {
					if pgerr.Code == pgUniqueViolation {
						s.error(w, r, http.StatusBadRequest, errorUniqueViolation)
						return
					}
				}
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			s.respond(w, r, http.StatusAccepted, company)
			return
		}

		s.respond(w, r, http.StatusCreated, u)
	}
}

// handleAccountDelete ...
func (s *Server) handleAccountDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(ctxKeyUser).(*model.User)
		if !ok {
			s.error(w, r, http.StatusInternalServerError, errorConvertation)
			return
		}

		err := s.store.User().Delete(user)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

// handleAuth ...
func (s *Server) handleAuth() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		user, err := s.store.User().FindByLogin(req.Login)
		if err != nil {
			if err == sql.ErrNoRows {
				s.error(w, r, http.StatusUnauthorized, errorIncorrectLoginOrPassword)
				return
			}
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if !service.EncryptCompare(user.Password, req.Password) {
			s.error(w, r, http.StatusUnauthorized, errorIncorrectLoginOrPassword)
			return
		}

		jwtClaim := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
		}
		tk := &model.Token{UserID: user.ID, StandardClaims: jwtClaim}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(s.secretKey))

		w.Header().Add("Authorization", tokenString)
		s.respond(w, r, http.StatusOK, user)
	}
}

//------------USER------------
// handleUserGet ...
func (s *Server) handleUserGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(ctxKeyUser).(*model.User)
		if !ok {
			s.error(w, r, http.StatusInternalServerError, errorConvertation)
			return
		}

		s.respond(w, r, http.StatusOK, user)
	}
}

// handleUserEdit ...
func (s *Server) handleUserEdit() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		user, ok := r.Context().Value(ctxKeyUser).(*model.User)
		if !ok {
			s.error(w, r, http.StatusInternalServerError, errorConvertation)
			return
		}

		if req.Login == "" && req.Password == "" {
			s.error(w, r, http.StatusBadRequest, errorBadRequest)
			return
		}
		if req.Login != "" {
			user.Login = req.Login
		}
		if req.Password != "" {
			user.Password = req.Password
		}

		err := s.store.User().Update(user)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, user)
	}
}

//------------COMPANY------------
// handleCompanyGet ...
func (s *Server) handleCompanyGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		company, ok := r.Context().Value(ctxKeyUser).(*model.Company)
		if !ok {
			s.error(w, r, http.StatusInternalServerError, errorConvertation)
			return
		}

		s.respond(w, r, http.StatusOK, company)
	}
}

// handleCompanyEdit ...
func (s *Server) handleCompanyEdit() http.HandlerFunc {
	type request struct {
		BIN     string `json:"bin"`
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		company, ok := r.Context().Value(ctxKeyUser).(*model.Company)
		if !ok {
			s.error(w, r, http.StatusInternalServerError, errorConvertation)
			return
		}

		company.BIN = req.BIN
		company.Name = req.Name
		company.Address = req.Address

		err := s.store.Company().Update(company)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, company)
	}
}

// handleCompanyLevelCreate ...
func (s *Server) handleLevelCreate() http.HandlerFunc {
	type request struct {
		Expreience int64 `json:"experience"`
		Level      int   `json:"level"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		company, ok := r.Context().Value(ctxKeyUser).(*model.Company)
		if !ok {
			s.error(w, r, http.StatusInternalServerError, errorConvertation)
			return
		}

		companyLevel := &model.CompanyLevel{
			Company:    company,
			Experience: req.Expreience,
			Level:      req.Level,
		}

		err := s.store.CompanyLevel().Save(companyLevel)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusCreated, companyLevel)
	}
}

// handleCompanyLevelGet ...
func (s *Server) handleLevelGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		company, ok := r.Context().Value(ctxKeyUser).(*model.Company)
		if !ok {
			s.error(w, r, http.StatusInternalServerError, errorConvertation)
			return
		}

		companyLevels, err := s.store.CompanyLevel().FindByCompany(company)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, companyLevels)
	}
}

func (s *Server) handleCompanyTransactionCreate() http.HandlerFunc {
	type request struct {
		UserID     int64 `json:"user_id"`
		Experience int64 `json:"experience"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		company, ok := r.Context().Value(ctxKeyUser).(*model.Company)
		if !ok {
			s.error(w, r, http.StatusInternalServerError, errorConvertation)
			return
		}

		user, err := s.store.User().FindByID(req.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				s.error(w, r, http.StatusNotFound, errorNotFound)
				return
			}
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		transaction := &model.Transaction{
			User:       user,
			Company:    company,
			Experience: req.Experience,
		}

		err = s.store.Transaction().Save(transaction)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusCreated, transaction)
	}
}

// handleUserTransactionGet ...
func (s *Server) handleUserTransactionGet() http.HandlerFunc {
	type request struct {
		CompanyID int64 `json:"company_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		user, ok := r.Context().Value(ctxKeyUser).(*model.User)
		if !ok {
			s.error(w, r, http.StatusInternalServerError, errorConvertation)
			return
		}

		company, err := s.store.Company().FindByID(req.CompanyID)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		transactions, err := s.store.Transaction().FindByUserAndCompany(user, company)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, transactions)
	}
}
