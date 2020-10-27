package controllers

import (
	"github.com/labstack/echo"
)

func InitScores() {

}

// Controllers
// func GetTeamScore(c echo.Context) error {
// 	scores := new(models.Score)
// 	var err error
// 	fmt.Println(err)
// 	return c.JSON(http.StatusCreated, scores)
// }

func (s *Score) GetScores(c echo.Context) {

}

func (s *Score) GetScore(c echo.Context) {

}

func (s *Score) PostScore(c echo.Context) {

}

func (s *Score) UpdateScore(c echo.Context) {

}

func (s *Score) PatchScore(c echo.Context) {

}

func (s *Score) DeleteScore(c echo.Context) {

}

func ScoresRoutes() {
	// New Echo instance
	e := echo.New()

	e.GET("/sportsapi/v1/scores/:sport", GetScores)
	e.GET("/sportsapi/v1/scores/:sport/:team", GetScore)
	e.PUT("/sportsapi/v1/scores/:sport/:team", UpdateScore)
	e.PATCH("/sportsapi/v1/scores/:sport/:team", PatchScore)
	e.POST("/sportsapi/v1/scores", PostScore)
	e.DELETE("/sportsapi/v1/scores/:sport/:team", DeleteScore)

}
