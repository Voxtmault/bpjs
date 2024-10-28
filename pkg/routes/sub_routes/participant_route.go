package sub_routes

import "github.com/labstack/echo/v4"

func Participant(group *echo.Group) {

	participant := group.Group("/participant")
	participant.GET("/nik/:nik", nil)
	participant.GET("/card_number/:nik", nil)
}
