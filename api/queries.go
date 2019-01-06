package api

import (
	"errors"
	"github.com/graphql-go/graphql"
)

func (s *Service) createExerciseQuery(exerciseType *graphql.Object) graphql.ObjectConfig {
	return graphql.ObjectConfig{
		Name: "ExerciseQuery",
		Fields: graphql.Fields{
			// TODO добавить выдачу по идентификтаору категории и по идентификатору упражнения
			"list": &graphql.Field{
				Type: graphql.NewList(exerciseType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					exercises, err := s.exerciseRepository.GetAll()
					if err != nil {
						return nil, err
					}
					for _, sx := range exercises {
						cat, err := s.categoryRepository.GetCategory(sx.CategoryID)
						if err != nil {
							return nil, err
						}
						sx.Category = *cat
					}
					return exercises, nil
				},
			},
		},
	}
}

func (s *Service) createCategoryQuery(categoryType *graphql.Object) graphql.ObjectConfig {
	return graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{

			"category": &graphql.Field{
				Type:        categoryType,
				Description: "Get category by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if !ok {
						return nil, errors.New("cant get category id")
					}
					cat, err := s.categoryRepository.GetCategory(int64(id))
					if err != nil {
						return nil, err
					}
					return cat, nil
				},
			},

			"list": &graphql.Field{
				Type:        graphql.NewList(categoryType),
				Description: "Get category list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					categories, err := s.categoryRepository.GetAllCategories()
					if err != nil {
						return nil, err
					}
					return categories, nil
				},
			},
		},
	}
}

func (s *Service) createSetQuery(setsType *graphql.Object) graphql.ObjectConfig {
	return graphql.ObjectConfig{
		Name: "Sets",
		Fields: graphql.Fields{
			"list": &graphql.Field{
				Type: graphql.NewList(setsType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					sets := s.setRepository.GetAll()
					resp := s.setRepository.GetOrderedByDate(sets)
					return resp, nil
				},
			},
		},
	}
}
