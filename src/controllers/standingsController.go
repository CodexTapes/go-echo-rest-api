package controllers

import (
	"fmt"
	"net/http"

	"go-echo-swagger-sports-api/src/models"

	"github.com/labstack/echo"
)

// Controllers
func InitStandings(c echo.Context) error {
	standings := new(models.Standings)
	var err error
	fmt.Println(err)
	return c.JSON(http.StatusCreated, team)
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

	e.GET("/sportsapi/v1/standings/:sport", GetStandings)
	e.GET("/sportsapi/v1/standings/:sport/:team", GetStanding)
	e.PUT("/sportsapi/v1/standings/:sport/:team", UpdateTeamStandings)
	e.POST("/sportsapi/v1/standings/", CreateTeamStandings)
	e.DELETE("/sportsapi/v1/standings/:sport/:team", DeleteTeamStandings)

}
