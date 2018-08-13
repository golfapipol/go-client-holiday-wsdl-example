package soap_test

import (
	. "holiday"
	. "holiday/soap"
	"testing"
)

func Test_RequestToWDSL_Input_HolidayRequest_Should_Be_Envelope_With_GetHolidaysAvailableResponse(t *testing.T) {
	expectedHolidays := 34
	holidayRequest := HolidayRequest{
		CountryCode: "UnitedStates",
	}
	actualHolidayResponse := RequestToWDSL(holidayRequest)
	if expectedHolidays != len(actualHolidayResponse.Holiday) {
		t.Errorf("expected \n%d\nbut it got\n%d\n%v", expectedHolidays, len(actualHolidayResponse.Holiday), actualHolidayResponse)
	}
}
