package repeat

import (
	"database/sql"
	"go-fitness/models"
)

type Service struct {
	db *sql.DB
}

func NewRepeatService(db *sql.DB) *Service {
	return &Service{db: db}
}


func (s *Service) GetRepeats(id int64) []models.Repeat {
	resp := make([]models.Repeat, 0)
	repsRow, err := s.db.Query("SELECT id, weight, count FROM repeats WHERE set_id = ?", id)
	if err != nil {
		panic(err)
	}

	for repsRow.Next() {
		s := &models.Repeat{}
		err = repsRow.Scan(&s.ID, &s.Weight, &s.Count)
		if err != nil {
			panic(err)
		}

		resp = append(resp, *s)
	}

	return resp

}
