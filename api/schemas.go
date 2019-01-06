package api

import "github.com/graphql-go/graphql"

func (s *Service) CategorySchema() (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				s.createCategoryQuery(
					s.createCategoryType(),
				),
			),
		},
	)
}

func (s *Service) ExercisesSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(
			s.createExerciseQuery(
				s.createExerciseType(
					s.createCategoryType(),
				),
			),
		),
	})
}

func (s *Service) SetSchema() (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				s.createSetQuery(
					s.createOrderedSetsType(
						s.createSetType(
							s.createRepeatType(),
							s.createExerciseType(
								s.createCategoryType(),
							),
						),
					),
				),
			),
		},
	)
}
