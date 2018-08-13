package api_test

import (
	"encoding/xml"
	. "holiday/api"
	"testing"
)

func Test_decodingXML_HolidayServiceWSDL_Should_Be_Formatted_In_Struct(t *testing.T) {
	var actualXML Envelope
	expectedHolidayServiceXML := Envelope{
		Soap: &SoapBody{
			GetHolidaysAvailableResponse: &GetHolidaysAvailableResponse{
				GetHolidaysAvailableResult: &GetHolidaysAvailableResult{
					HolidayCode: &[]HolidayCode{
						HolidayCode{Code: "NEW-YEARS-DAY-ACTUAL", Description: "New Year's Day"},
						HolidayCode{Code: "NEW-YEARS-DAY-OBSERVED", Description: "New Year's Day"},
					},
				},
			},
		},
	}
	wsdl := `
<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <soap:Body>
        <GetHolidaysAvailableResponse xmlns="http://www.holidaywebservice.com/HolidayService_v2/">
            <GetHolidaysAvailableResult>
                <HolidayCode>
                    <Code>NEW-YEARS-DAY-ACTUAL</Code>
                    <Description>New Year's Day</Description>
                </HolidayCode>
                <HolidayCode>
                    <Code>NEW-YEARS-DAY-OBSERVED</Code>
                    <Description>New Year's Day</Description>
                </HolidayCode>
            </GetHolidaysAvailableResult>
        </GetHolidaysAvailableResponse>
    </soap:Body>
</soap:Envelope>`
	xml.Unmarshal([]byte(wsdl), &actualXML)
	actualHolidayCodes := *actualXML.Soap.GetHolidaysAvailableResponse.GetHolidaysAvailableResult.HolidayCode
	for index, expectedHolidayCode := range *expectedHolidayServiceXML.Soap.GetHolidaysAvailableResponse.GetHolidaysAvailableResult.HolidayCode {
		actualHolidayCode := actualHolidayCodes[index]
		if expectedHolidayCode != actualHolidayCode {
			t.Errorf("expected index: %d HolidayCode is %v but it got %v", index, expectedHolidayCode, actualHolidayCode)
		}
	}

}

func Test_encodingXML_GetHolidaysAvailableResponse_Should_Be_Formatted_In_XML(t *testing.T) {
	GetHolidaysAvailableResponse := GetHolidaysAvailableResponse{
		GetHolidaysAvailableResult: &GetHolidaysAvailableResult{
			HolidayCode: &[]HolidayCode{
				HolidayCode{Code: "NEW-YEARS-DAY-ACTUAL", Description: "New Year's Day"},
				HolidayCode{Code: "NEW-YEARS-DAY-OBSERVED", Description: "New Year's Day"},
			},
		},
	}
	expectedResponse := `<GetHolidaysAvailableResponse xmlns="http://www.holidaywebservice.com/HolidayService_v2/"><GetHolidaysAvailableResult><HolidayCode><Code>NEW-YEARS-DAY-ACTUAL</Code><Description>New Year&#39;s Day</Description></HolidayCode><HolidayCode><Code>NEW-YEARS-DAY-OBSERVED</Code><Description>New Year&#39;s Day</Description></HolidayCode></GetHolidaysAvailableResult></GetHolidaysAvailableResponse>`
	actualResponse, _ := xml.Marshal(GetHolidaysAvailableResponse)
	if expectedResponse != string(actualResponse) {
		t.Errorf("expected \n%s but it got \n%s", expectedResponse, actualResponse)
	}

}
