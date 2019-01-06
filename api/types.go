package api

import (
	"github.com/graphql-go/graphql"
	"go-fitness/models"
)

func (s *Service) createCategoryType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "category",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func (s *Service) createExerciseType(categoryType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Exercise",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"category": &graphql.Field{
				Type: categoryType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					exs := p.Source.(*models.Exercise)
					cat, err := s.categoryRepository.GetCategory(exs.CategoryID)
					if err != nil {
						return nil, err
					}
					return cat, nil
				},
			},
		},
	})
}

func (s *Service) createRepeatType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Repeat",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
				Description: "id",
			},
			"weight": &graphql.Field{
				Type: graphql.Int,
				Description: "weight",
			},
			"count": &graphql.Field{
				Type: graphql.Int,
				Description: "count",
			},
		},
	})
}

func (s *Service) createSetType(repeatType *graphql.Object, exerciseType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Set",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "id",
			},
			"date": &graphql.Field{
				Type:        graphql.DateTime,
				Description: "date",
			},
			"exercise": &graphql.Field{
				Type: exerciseType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					set, _ := p.Source.(*models.Set)
					ex := s.exerciseRepository.GetExercise(set.ExerciseID)
					return ex, nil
				},
			},
			"repeats": &graphql.Field{
				Type: graphql.NewList(repeatType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					set, _ := p.Source.(*models.Set)
					resp := s.repeatRepository.GetRepeats(set.ID)
					return resp, nil
				},
			},
		},
	})
}

func (s *Service) createOrderedSetsType(setType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "SetsByDate",
		Fields: graphql.Fields{
			"date": &graphql.Field{Type: graphql.DateTime},
			"data": &graphql.Field{Type: graphql.NewList(setType)},
		},
	})
}
