package api

import "go-fitness/models"

// SetRepository интерфейс сервиса подходов
type SetRepository interface {
	GetAll() models.Sets
	GetOrderedByDate(sets []*models.Set) []models.SetsOrdered
}

// RepeatRepository интерфейс сервиса повторений
type RepeatRepository interface {
	GetRepeats(id int64) []models.Repeat
}


// ExerciseRepository интерфейс сервиса упражнений
type ExerciseRepository interface {
	GetAll() (models.Exercises, error)
	GetExercise(id int64) *models.Exercise
}


// CategoryRepository интерфейс сервиса категорий упражений
type CategoryRepository interface {
	GetAllCategories() (models.Categories, error)
	GetCategory(id int64) (*models.Category, error)
}
