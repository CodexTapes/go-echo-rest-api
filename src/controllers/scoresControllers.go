package controllers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/swag"
)

func InitScores() {

}


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

	scores := e.Group("/echosportsapi/v1/scores")

	scores.Use(middleware.LoggerWithConfig(middleware.LoggerConfig {
		format: `[$(time_rfc3339)] $(method) $(uri) $(status)` + "\n"
	}))

	scores.GET("/:sport", GetScores)
	scores.GET("/:sport/:team", GetScore)
	scores.PUT("/:sport/:team", UpdateScore)
	scores.PATCH("/:sport/:team", PatchScore)
	scores.POST("/", PostScore)
	scores.DELETE("/:sport/:team", DeleteScore)

}
