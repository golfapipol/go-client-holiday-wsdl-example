package api

import (
	. "holiday"
	"holiday/soap"
	"net/http"

	"github.com/labstack/echo"
)

func SetupRoute() *echo.Echo {
	router := echo.New()
	router.POST("/api/v1/holidays", GetHolidaysAvailableHandler)
	return router
}

func GetHolidaysAvailableHandler(context echo.Context) error {
	holidayRequest := new(HolidayRequest)
	if err := context.Bind(holidayRequest); err != nil {
		return err
	}
	holidayResponse := soap.RequestToWDSL(*holidayRequest)
	return context.JSON(http.StatusOK, holidayResponse)
}
