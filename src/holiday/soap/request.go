package soap

import (
	"bytes"
	"encoding/xml"
	. "holiday"
	"io/ioutil"
	"net/http"
)

func RequestToWDSL(request HolidayRequest) HolidayResponse {
	var envelope Envelope
	response, _ := http.Post("http://www.holidaywebservice.com/HolidayService_v2/HolidayService2.asmx?wsdl", "text/xml", bytes.NewBuffer([]byte(request.ToXML())))
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	xml.Unmarshal(body, &envelope)
	return envelope.Soap.GetHolidaysAvailableResponse.ToJSON()
}
