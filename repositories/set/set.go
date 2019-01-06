package set

import (
	"database/sql"
	"go-fitness/models"
)

type Service struct {
	db *sql.DB
}

func NewSetService(db *sql.DB) *Service {
	return &Service{db:db}
}


func (s *Service) GetAll() models.Sets {
	sets := make(models.Sets, 0)
	setRows, err := s.db.Query("SELECT id, date, exercise_id FROM `sets` ORDER BY date")
	if err != nil {
		panic(err)
	}

	for setRows.Next() {
		s := &models.Set{}
		err = setRows.Scan(&s.ID, &s.Date, &s.ExerciseID)
		if err != nil {
			panic(err)
		}

		sets = append(sets, s)
	}

	return sets
}

func (s *Service) GetOrderedByDate(sets []*models.Set) []models.SetsOrdered {
	resp := make([]models.SetsOrdered, 0)
	for _, s := range sets {
		indx, ok := checkExists(resp, s)
		if ok {
			resp[indx].Data = append(resp[indx].Data, s)
		} else {
			resp = append(resp, newSetsByDate(s))
		}
	}
	return resp
}

func newSetsByDate(s *models.Set) models.SetsOrdered {
	os := models.SetsOrdered{
		Date: s.Date,
	}
	os.Data = append(os.Data, s)

	return os
}

func checkExists(os []models.SetsOrdered, s *models.Set) (int, bool) {
	for i, is := range os {
		if s.Date == is.Date {
			return i, true
		}
	}
	return 0, false
}
