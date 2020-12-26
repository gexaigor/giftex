package apiserver

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gexaigor/MyRestAPI/store/sqlstore"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DataBaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.NewStore(db)

	s, err := newServer(store, config.SecretKey)
	if err != nil {
		log.Fatal(err)
	}

	s.logger.Info("Start ListenAndServe")

	return http.ListenAndServe(config.BindAddr, s)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
