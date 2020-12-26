package sqlstore

import (
	"time"

	"github.com/gexaigor/MyRestAPI/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Save ...
func (r *UserRepository) Save(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeSave(); err != nil {
		return err
	}

	user.CreatedOn = time.Now()

	return r.store.db.QueryRow(
		"INSERT INTO usr (login, email, password, role, created_on) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.Login,
		user.Email,
		user.Password,
		user.Role.String(),
		user.CreatedOn,
	).Scan(&user.ID)
}

// Update ...
func (r *UserRepository) Update(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeSave(); err != nil {
		return err
	}

	if _, err := r.store.db.Exec(
		"UPDATE usr SET login = $1, password = $2 WHERE id = $3",
		user.Login,
		user.Password,
		user.ID,
	); err != nil {
		return err
	}

	return nil
}

// Delete ...
func (r *UserRepository) Delete(user *model.User) error {
	if _, err := r.store.db.Exec(
		"DELETE FROM usr WHERE id = $1",
		user.ID,
	); err != nil {
		return err
	}

	return nil
}

// FindByID ...
func (r *UserRepository) FindByID(id int64) (*model.User, error) {
	user := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, login, email, role, created_on FROM usr WHERE id = $1",
		id,
	).Scan(
		&user.ID,
		&user.Login,
		&user.Email,
		&user.Role,
		&user.CreatedOn,
	); err != nil {
		return nil, err
	}

	return user, nil
}

// FindByLogin ...
func (r *UserRepository) FindByLogin(login string) (*model.User, error) {
	user := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, login, email, password, role, created_on FROM usr WHERE login = $1",
		login,
	).Scan(
		&user.ID,
		&user.Login,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedOn,
	); err != nil {
		return nil, err
	}

	return user, nil
}
