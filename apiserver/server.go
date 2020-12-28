package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gexaigor/MyRestAPI/store"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	ctxKeyUser ctxKey = iota
)

type ctxKey int8

// Server ...
type Server struct {
	router    *mux.Router
	logger    *zap.Logger
	store     store.Store
	secretKey string
}

func newServer(store store.Store, secretKey string) (*Server, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	s := &Server{
		router:    mux.NewRouter(),
		logger:    logger,
		store:     store,
		secretKey: secretKey,
	}

	s.configureRouter()

	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.Use(s.jwtMiddleware)
	s.router.HandleFunc("/auth", s.handleAuth()).Methods("POST")

	s.router.HandleFunc("/account", s.handleAccountCreate()).Methods("POST")
	s.router.HandleFunc("/account", s.handleAccountDelete()).Methods("DELETE")

	user := s.router.PathPrefix("/user").Subrouter()
	user.Use(s.userMiddleware)
	user.HandleFunc("", s.handleUserGet()).Methods("GET")
	user.HandleFunc("", s.handleUserEdit()).Methods("PUT")

	userTransaction := user.PathPrefix("/transaction").Subrouter()
	userTransaction.HandleFunc("", s.handleUserTransactionGet()).Methods("GET")

	company := s.router.PathPrefix("/company").Subrouter()
	company.Use(s.companyMiddleware)
	company.HandleFunc("", s.handleCompanyGet()).Methods("GET")
	company.HandleFunc("", s.handleCompanyEdit()).Methods("PUT")

	level := company.PathPrefix("/level").Subrouter()
	level.HandleFunc("", s.handleLevelCreate()).Methods("POST")
	level.HandleFunc("", s.handleLevelGet()).Methods("GET")

	companyTransaction := company.PathPrefix("/transaction").Subrouter()
	companyTransaction.HandleFunc("", s.handleCompanyTransactionCreate()).Methods("POST")
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}
