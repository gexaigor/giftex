package sqlstore

import "github.com/gexaigor/MyRestAPI/model"

// TransactionRepository ...
type TransactionRepository struct {
	store *Store
}

// Save ...
func (r *TransactionRepository) Save(transaction *model.Transaction) error {
	if err := transaction.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO transactions (usr_id, company_id, experience) VALUES ($1, $2, $3) RETURNING id",
		transaction.User.ID,
		transaction.Company.ID,
		transaction.Experience,
	).Scan(&transaction.ID)
}

// FindByUserAndCompany ...
func (r *TransactionRepository) FindByUserAndCompany(user *model.User, company *model.Company) (*[]model.Transaction, error) {
	transactions := []model.Transaction{}

	rows, err := r.store.db.Query(
		"SELECT id, experience FROM transactions WHERE usr_id = $1 AND company_id = $2",
		user.ID,
		company.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := model.Transaction{
			User:    user,
			Company: company,
		}
		err := rows.Scan(&t.ID, &t.Experience)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return &transactions, nil
}
