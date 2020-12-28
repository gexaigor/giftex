package sqlstore

import (
	"github.com/gexaigor/MyRestAPI/model"
)

// CompanyLevelRepository ...
type CompanyLevelRepository struct {
	store *Store
}

// Save ...
func (r *CompanyLevelRepository) Save(companylevel *model.CompanyLevel) error {
	if err := companylevel.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO company_levels (company_id, experience, level, description) VALUES ($1, $2, $3, $4) RETURNING id",
		companylevel.Company.ID,
		companylevel.Experience,
		companylevel.Level,
		companylevel.Description,
	).Scan(&companylevel.ID)
}

// FindByCompany ...
func (r *CompanyLevelRepository) FindByCompany(company *model.Company) (*[]model.CompanyLevel, error) {
	companyLevels := []model.CompanyLevel{}

	rows, err := r.store.db.Query(
		"SELECT id, experience, level FROM company_levels WHERE company_id = $1 ORDER BY level",
		company.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		cl := model.CompanyLevel{
			Company: company,
		}
		err := rows.Scan(&cl.ID, &cl.Experience, &cl.Level)
		if err != nil {
			return nil, err
		}
		companyLevels = append(companyLevels, cl)
	}

	return &companyLevels, nil
}
