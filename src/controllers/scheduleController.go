package controllers

import (
	"fmt"
	"net/http"

	"go-echo-swagger-sports-api/src/models"

	"github.com/labstack/echo"
)

func InitSchedule() {

}

// Controllers
func GetSchedule(c echo.Context) error {
	team := new(models.Team)
	var err error
	fmt.Println(err)
	return c.JSON(http.StatusCreated, team)
}

func (sch *Schedule) GetSchedule(c echo.Context) {

}

func (sch *Schedule) PostSchedule(c echo.Context) {

}

func (sch *Schedule) UpdateSchedule(c echo.Context) {

}

func (sch *Schedule) PatchSchedule(c echo.Context) {

}

func (sch *Schedule) DeleteSchedule(c echo.Context) {

}

func ScheduleRoutes() {
	// New Echo instance
	e := echo.New()

	e.GET("/sportsapi/v1/schedule/:sport", GetSchedule)
	e.GET("/sportsapi/v1/schedule/:sport/:team", GetTeamSchedule)
	e.PUT("/sportsapi/v1/schedule/:sport/:team", UpdateTeamSchedule)
	e.POST("/sportsapi/v1/schedule", CreateTeamSchedule)
	e.DELETE("/sportsapi/v1/schedule/:sport/:team", DeleteTeamSchedule)
}
