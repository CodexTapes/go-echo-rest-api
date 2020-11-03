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

	schedule := e.Group("/echosportsapi/v1/schedule")

	schedule.Use(middleware.LoggerWithConfig(middleware.LoggerConfig {
		format: `[$(time_rfc3339)] $(method) $(uri) $(status)` + "\n"
	}))

	schedule.GET("/:sport", GetSchedule)
	schedule.GET("/:sport/:team", GetTeamSchedule)
	schedule.PUT("/:sport/:team", UpdateTeamSchedule)
	schedule.POST("/", CreateTeamSchedule)
	schedule.DELETE("/:sport/:team", DeleteTeamSchedule)
}
