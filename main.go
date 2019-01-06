package main

import (
	"github.com/labstack/echo"
	"go-fitness/api"
	"go-fitness/database"
	"go-fitness/handlers"
	"go-fitness/repositories/category"
	"go-fitness/repositories/exercise"
	"go-fitness/repositories/repeat"
	"go-fitness/repositories/set"
)



func main() {
	db, err := database.NewDB("file:app.db??cache=shared&mode=rwc")
	if err != nil {
		panic(err)
	}

	catSrv := category.NewCategoryService(db)
	exSrv := exercise.NewExerciseService(db)
	setSrv := set.NewSetService(db)
	repSrv := repeat.NewRepeatService(db)

	apiService := api.New(catSrv, exSrv, repSrv, setSrv)

	e := echo.New()

	e.GET("/category", handlers.Category(apiService))
	e.GET("/exercise", handlers.Exercise(apiService))
	e.GET("/set", handlers.Set(apiService))

	e.Logger.Debug(e.Start(":1323"))
}



