package soap

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	. "holiday"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func RequestToWDSL(holidayRequest HolidayRequest) HolidayResponse {
	var envelope Envelope
	request, _ := http.NewRequest("POST", "http://www.holidaywebservice.com/HolidayService_v2/HolidayService2.asmx?wsdl", bytes.NewBuffer([]byte(holidayRequest.ToXML())))
	request.Header.Set(echo.HeaderContentType, echo.MIMETextXML)
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*30)
	request = request.WithContext(ctx)
	response, err := http.DefaultClient.Do(request.WithContext(ctx))
	if err != nil {
		fmt.Println(err.Error())
		return HolidayResponse{}
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	xml.Unmarshal(body, &envelope)
	return envelope.Soap.GetHolidaysAvailableResponse.ToJSON()
}
