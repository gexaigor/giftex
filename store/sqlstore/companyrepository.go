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

// FindAll ...
func (r *CompanyRepository) FindAll(page int, limit int) ([]model.Company, error) {
	companys := []model.Company{}

	offset := limit * (page - 1)

	rows, err := r.store.db.Query(
		"SELECT c.id, c.bin, c.name, c.address, u.id, u.login, u.email, u.role, u.created_on FROM company AS c "+
			"LEFT JOIN usr AS u ON u.id = c.usr_id "+
			"ORDER BY c.id LIMIT $1 OFFSET $2",
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := &model.Company{}
		u := &model.User{}
		err := rows.Scan(&c.ID, &c.BIN, &c.Name, &c.Address, &u.ID, &u.Login, &u.Email, &u.Role, &u.CreatedOn)
		if err != nil {
			return nil, err
		}
		c.User = u
		companys = append(companys, *c)
	}

	return companys, nil
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
