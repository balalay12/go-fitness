package api

import (
	"github.com/graphql-go/graphql"
	"log"
)

type Service struct {
	categoryRepository CategoryRepository
	exerciseRepository ExerciseRepository
	repeatRepository   RepeatRepository
	setRepository      SetRepository
}

func New(cat CategoryRepository, ex ExerciseRepository, rep RepeatRepository, set SetRepository) *Service {
	return &Service{
		categoryRepository: cat,
		exerciseRepository: ex,
		repeatRepository: rep,
		setRepository: set,
	}
}

func (s *Service) ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Fatalf("errors: %v", result.Errors)
	}
	return result
}
