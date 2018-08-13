package api

import (
	"encoding/xml"
)

type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Soap    *SoapBody
}
type SoapBody struct {
	XMLName                      xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	GetHolidaysAvailableResponse *GetHolidaysAvailableResponse
}
type GetHolidaysAvailableResponse struct {
	XMLName                    xml.Name `xml:"http://www.holidaywebservice.com/HolidayService_v2/ GetHolidaysAvailableResponse"`
	GetHolidaysAvailableResult *GetHolidaysAvailableResult
}
type GetHolidaysAvailableResult struct {
	HolidayCode *[]HolidayCode
}
type HolidayCode struct {
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}

func (response GetHolidaysAvailableResponse) ToJSON() HolidayResponse {
	holidays := make([]Holiday, len(*response.GetHolidaysAvailableResult.HolidayCode))
	for index := range holidays {
		holidays[index] = (*response.GetHolidaysAvailableResult.HolidayCode)[index].ToJSON()
	}
	return HolidayResponse{
		Holiday: holidays,
	}
}

func (hc HolidayCode) ToJSON() Holiday {
	return Holiday{
		Code:        hc.Code,
		Description: hc.Description,
	}
}
