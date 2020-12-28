package store

import (
	"github.com/gexaigor/MyRestAPI/model"
)

// UserRepository ...
type UserRepository interface {
	Save(user *model.User) error
	Delete(user *model.User) error
	Update(user *model.User) error
	FindByID(id int64) (*model.User, error)
	FindByLogin(login string) (*model.User, error)
}

// CompanyRepository ...
type CompanyRepository interface {
	Save(company *model.Company) error
	Update(company *model.Company) error
	FindAll(page int, limit int) ([]model.Company, error)
	FindByID(id int64) (*model.Company, error)
	FindByUser(user *model.User) (*model.Company, error)
}

// CompanyLevelRepository ...
type CompanyLevelRepository interface {
	Save(companyLevel *model.CompanyLevel) error
	FindByCompany(company *model.Company) (*[]model.CompanyLevel, error)
}

// TransactionRepository ...
type TransactionRepository interface {
	Save(transaction *model.Transaction) error
	FindByUserAndCompany(user *model.User, company *model.Company) (*[]model.Transaction, error)
}
