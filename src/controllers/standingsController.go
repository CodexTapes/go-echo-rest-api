package controllers

import (
	"fmt"
	"net/http"

	"go-echo-swagger-sports-api/src/models"

	"github.com/labstack/echo"
)

// Controllers
func InitStandings(c echo.Context) error {
	
}

func (st *Standings) GetStandings(c echo.Context) {

}

func (st *Standings) GetStanding(c echo.Context) {

}

func (st *Standings) GetTeamStandings(c echo.Context) {

}

func (st *Standings) CreateTeamStandings(c echo.Context) {

}

func (st *Standings) UpdateStandings(c echo.Context) {

}

func (st *Standings) DeleteTeamStandings(c echo.Context) {

}

func StandingsRoutes() {
	// New Echo instance
	e := echo.New()

	standings := e.Group("/echosportsapi/v1/standings")

	standings.Use(middleware.LoggerWithConfig(middleware.LoggerConfig {
		format: `[$(time_rfc3339)] $(method) $(uri) $(status)` + "\n"
	}))

	standings.GET("/:sport", GetStandings)
	standings.GET("/:sport/:team", GetStanding)
	standings.PUT("/:sport/:team", UpdateTeamStandings)
	standings.POST("/", CreateTeamStandings)
	standings.DELETE("/:sport/:team", DeleteTeamStandings)

}
