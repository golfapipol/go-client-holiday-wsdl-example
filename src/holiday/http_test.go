package holiday_test

import (
	"bytes"
	"encoding/xml"
	. "holiday"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_HTTPRequest_Input_GetHolidaysAvailable_Should_Be_GetHolidaysAvailableResponse(t *testing.T) {
	var actualEnvelope Envelope
	expectedHolidayCodes := 34
	requestXML := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:hs="http://www.holidaywebservice.com/HolidayService_v2/">
	<soapenv:Body>
		<hs:GetHolidaysAvailable>
			<hs:countryCode>UnitedStates</hs:countryCode>
		</hs:GetHolidaysAvailable>
	</soapenv:Body>
</soapenv:Envelope>`
	response, _ := http.Post("http://www.holidaywebservice.com/HolidayService_v2/HolidayService2.asmx?wsdl", "text/xml", bytes.NewBuffer([]byte(requestXML)))
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	xml.Unmarshal(body, &actualEnvelope)

	actualHolidayCodes := len(*actualEnvelope.Soap.GetHolidaysAvailableResponse.GetHolidaysAvailableResult.HolidayCode)
	if expectedHolidayCodes != actualHolidayCodes {
		t.Errorf("expected %d but it got %d", expectedHolidayCodes, actualHolidayCodes)
	}

}
