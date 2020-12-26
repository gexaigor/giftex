package sqlstore

import (
	"database/sql"

	"github.com/gexaigor/MyRestAPI/store"
)

// Store ...
type Store struct {
	db                     *sql.DB
	userRepository         *UserRepository
	companyRepository      *CompanyRepository
	companyLevelRepository *CompanyLevelRepository
	transactionRepository  *TransactionRepository
}

// NewStore ...
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

// Company ...
func (s *Store) Company() store.CompanyRepository {
	if s.companyRepository != nil {
		return s.companyRepository
	}

	s.companyRepository = &CompanyRepository{
		store: s,
	}

	return s.companyRepository
}

// CompanyLevel ...
func (s *Store) CompanyLevel() store.CompanyLevelRepository {
	if s.companyLevelRepository != nil {
		return s.companyLevelRepository
	}

	s.companyLevelRepository = &CompanyLevelRepository{
		store: s,
	}

	return s.companyLevelRepository
}

// Transaction ...
func (s *Store) Transaction() store.TransactionRepository {
	if s.transactionRepository != nil {
		return s.transactionRepository
	}

	s.transactionRepository = &TransactionRepository{
		store: s,
	}

	return s.transactionRepository
}
