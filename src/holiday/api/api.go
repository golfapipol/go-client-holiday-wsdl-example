package api

import (
	"encoding/xml"
	"fmt"
	"holiday"
	"net/http"

	"github.com/labstack/echo"
)

type HolidayRequest struct {
	CountryCode string `json:"countryCode"`
}
type HolidayResponse struct {
	Holiday []Holiday `json:"holidays"`
}
type Holiday struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

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

	holidayResponse := HolidayResponse{}
	return context.JSON(http.StatusOK, holidayResponse)
}

func (request HolidayRequest) ToXML() string {
	getHolidaysAvailable := holiday.GetHolidaysAvailable{
		CountryCode: request.CountryCode,
	}
	xmlData, _ := xml.Marshal(getHolidaysAvailable)
	return fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:hs="http://www.holidaywebservice.com/HolidayService_v2/"><soapenv:Body>%s</soapenv:Body></soapenv:Envelope>`, xmlData)
}
