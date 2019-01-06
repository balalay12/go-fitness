package exercise

import (
	"database/sql"
	"go-fitness/models"
)

type Service struct {
	db *sql.DB
}

func NewExerciseService(db *sql.DB) *Service {
	return &Service{db:db}
}

func (s *Service) GetAll() (models.Exercises, error) {
	exercises := make(models.Exercises, 0)
	exerciseRow, err := s.db.Query("SELECT id, name, category_id FROM exercises")
	if err != nil {
		return nil, err
	}

	for exerciseRow.Next() {
		s := &models.Exercise{}
		err = exerciseRow.Scan(&s.ID, &s.Name, &s.CategoryID)
		if err != nil {
			return nil, err
		}

		exercises = append(exercises, s)
	}

	return exercises, nil
}

func(s *Service) GetExercise(id int64) *models.Exercise {
	exerciseRow, err := s.db.Query("SELECT id, name, category_id FROM exercises WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	ex := &models.Exercise{}
	for exerciseRow.Next() {
		err = exerciseRow.Scan(&ex.ID, &ex.Name, &ex.CategoryID)
		if err != nil {
			panic(err)
		}

	}

	return ex
}
