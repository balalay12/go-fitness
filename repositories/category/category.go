package category

import (
	"database/sql"
	"go-fitness/models"
)

type Service struct {
	db *sql.DB
}

func NewCategoryService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetAllCategories() (models.Categories, error) {
	categories := make(models.Categories, 0)
	catRows, err := s.db.Query("SELECT id, name FROM `categories`")
	if err != nil {
		return models.Categories{}, err
	}

	for catRows.Next() {
		c := &models.Category{}
		err = catRows.Scan(&c.ID, &c.Name)
		if err != nil {
			return models.Categories{}, err
		}

		categories = append(categories, c)
	}

	return categories, nil
}

func (s *Service) GetCategory(id int64) (*models.Category, error) {
	catRows, err := s.db.Query("SELECT id, name FROM `categories` WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	c := &models.Category{}
	for catRows.Next() {
		err = catRows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
