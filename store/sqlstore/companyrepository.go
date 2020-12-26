package sqlstore

import (
	"github.com/gexaigor/MyRestAPI/model"
)

// CompanyRepository ...
type CompanyRepository struct {
	store *Store
}

// Save ...
func (r *CompanyRepository) Save(company *model.Company) error {
	if err := company.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO company (usr_id, bin, name, address) VALUES ($1, $2, $3, $4) RETURNING id",
		company.User.ID,
		company.BIN,
		company.Name,
		company.Address,
	).Scan(&company.ID)
}

// Update ...
func (r *CompanyRepository) Update(company *model.Company) error {
	if err := company.Validate(); err != nil {
		return err
	}

	if _, err := r.store.db.Exec(
		"UPDATE company SET bin = $1, name = $2, address = $3 WHERE id = $4",
		company.BIN,
		company.Name,
		company.Address,
		company.ID,
	); err != nil {
		return err
	}

	return nil
}

// FindByID ...
func (r *CompanyRepository) FindByID(id int64) (*model.Company, error) {
	company := &model.Company{}
	var userID int64
	if err := r.store.db.QueryRow(
		"SELECT id, usr_id, bin, name, address FROM company WHERE id = $1",
		id,
	).Scan(
		&company.ID,
		&userID,
		&company.BIN,
		&company.Name,
		&company.Address,
	); err != nil {
		return nil, err
	}

	user, err := r.store.User().FindByID(userID)
	if err != nil {
		return nil, err
	}

	company.User = user

	return company, nil
}

// FindByUser ...
func (r *CompanyRepository) FindByUser(user *model.User) (*model.Company, error) {
	company := &model.Company{User: user}
	if err := r.store.db.QueryRow(
		"SELECT id, bin, name, address FROM company WHERE usr_id = $1",
		user.ID,
	).Scan(
		&company.ID,
		&company.BIN,
		&company.Name,
		&company.Address,
	); err != nil {
		return nil, err
	}

	return company, nil
}
