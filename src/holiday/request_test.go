package holiday_test

import (
	"encoding/xml"
	. "holiday"
	"testing"
)

func Test_encodingXML_GetHolidaysAvailable_Should_Be_Formatted_In_XML(t *testing.T) {
	expectedXML := `<hs:GetHolidaysAvailable><hs:countryCode>UnitedStates</hs:countryCode></hs:GetHolidaysAvailable>`
	getHolidaysAvailable := GetHolidaysAvailable{
		CountryCode: "UnitedStates",
	}
	actualXML, _ := xml.Marshal(getHolidaysAvailable)
	if string(actualXML) != expectedXML {
		t.Errorf("expected \n%s but it got \n%s", expectedXML, actualXML)
	}
}
