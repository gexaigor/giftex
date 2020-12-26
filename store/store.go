package store

// Store ...
type Store interface {
	User() UserRepository
	Company() CompanyRepository
	CompanyLevel() CompanyLevelRepository
	Transaction() TransactionRepository
}
