package soap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"holiday/api"
	"io/ioutil"
	"net/http"
)

func RequestToWDSL(request api.HolidayRequest) api.HolidayResponse {
	var envelope api.Envelope
	fmt.Println("RequestToWDSL\n", request.ToXML())
	response, _ := http.Post("http://www.holidaywebservice.com/HolidayService_v2/HolidayService2.asmx?wsdl", "text/xml", bytes.NewBuffer([]byte(request.ToXML())))
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	xml.Unmarshal(body, &envelope)
	fmt.Println("body", string(body))
	return envelope.Soap.GetHolidaysAvailableResponse.ToJSON()
}
