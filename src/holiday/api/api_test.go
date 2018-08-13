package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	. "holiday"
	. "holiday/api"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo"
)

func Test_GetHolidaysAvailableHandler_Input_UnitedStates_Should_Be_34_days(t *testing.T) {
	expectedHolidays := 34
	var actualHolidayResponse HolidayResponse
	requestData := []byte(`{"countryCode":"UnitedStates"}`)
	request := httptest.NewRequest("POST", "/api/v1/holidays", bytes.NewBuffer(requestData))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*30)
	recorder := httptest.NewRecorder()
	route := SetupRoute()
	route.ServeHTTP(recorder, request.WithContext(ctx))
	response := recorder.Result()
	json.NewDecoder(response.Body).Decode(&actualHolidayResponse)
	if expectedHolidays != len(actualHolidayResponse.Holiday) {
		t.Errorf("expected %d but it got %d", expectedHolidays, len(actualHolidayResponse.Holiday))
	}
}
