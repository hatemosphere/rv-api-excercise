package tests

import (
	"testing"
	"time"

	"github.com/bouk/monkey"
	"github.com/hatemosphere/rv-api-excercise/api/handlers/users"
	"github.com/stretchr/testify/assert"
)

func TestDateStringProcessorSuccess(t *testing.T) {
	expectedResult := time.Date(
		2019, 01, 01, 0, 0, 0, 0, time.UTC)
	result, ok := users.DateStringProcessor(users.DateLayout, "2019-01-01")
	assert.Equal(t, expectedResult, result, ok, "")
}

func TestDateDifferSucess(t *testing.T) {
	waybackNow := time.Date(
		2017, 01, 01, 0, 0, 00, 0, time.UTC)
	date := time.Date(
		2016, 01, 01, 0, 0, 0, 0, time.UTC)
	patch := monkey.Patch(time.Now, func() time.Time { return waybackNow })
    defer patch.Unpatch()
	expectedDiff := 366
	result := users.DateDiffer(date)
	assert.Equal(t, expectedDiff, result, "")
}

func TestBirhdayPrinterSucess(t *testing.T) {
	waybackNow := time.Date(
		2017, 01, 02, 0, 0, 00, 0, time.UTC)
	username := "john_doe"
	userBirthdayDate := time.Date(
		2017, 01, 01, 0, 0, 0, 0, time.UTC)
	patch := monkey.Patch(time.Now, func() time.Time { return waybackNow })
	defer patch.Unpatch()
	expectedResult := "Hello, john_doe! Your birthday is in 364 day(s)"
	result, ok := users.BirhdayPrinter(username, userBirthdayDate)
	assert.Equal(t, expectedResult, result, ok, "")
}
