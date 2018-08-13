package soap_test

import (
	"holiday/api"
	"testing"
)

func Test_HolidayRequestToXML_Should_Be_GetHolidaysAvailable(t *testing.T) {
	holidayResponse := api.HolidayRequest{
		CountryCode: "UnitedStates",
	}
	expectedXML := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:hs="http://www.holidaywebservice.com/HolidayService_v2/"><soapenv:Body><hs:GetHolidaysAvailable><hs:countryCode>UnitedStates</hs:countryCode></hs:GetHolidaysAvailable></soapenv:Body></soapenv:Envelope>`
	actualXML := holidayResponse.ToXML()
	if expectedXML != actualXML {
		t.Errorf("expected \n%s \nbut it got \n%s", expectedXML, actualXML)
	}
}
