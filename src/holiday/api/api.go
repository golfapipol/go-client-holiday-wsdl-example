package api

import (
	"fmt"
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
	fmt.Println(holidayRequest)
	holidayResponse := soap.RequestToWDSL(*holidayRequest)
	return context.JSON(http.StatusOK, holidayResponse)
}
