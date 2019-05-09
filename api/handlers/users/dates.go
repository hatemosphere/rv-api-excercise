package users

import (
	"fmt"
	"log"
	"time"
)

const (
	dateLayout = "2006-01-02"
)

// DateStringProcessor function converts string to time type
func DateStringProcessor(dateLayout, dateString string) (time.Time, bool) {
	date, err := time.Parse(dateLayout, dateString)
	if err != nil {
		log.Printf("Cannot parse date string: %v", err)
		return time.Time{}, false
	}
	return date, true
}

// DateDiffer function is used to validate date of birth
func DateDiffer(date time.Time) int {
	timeNow := time.Now()
	diff := timeNow.Sub(date)
	daysDiff := int(diff.Hours() / 24)
	return daysDiff
}

// BirthDayPrinter function prints a birthday message for username
func BirthDayPrinter(username string, userBirthdayDate time.Time) (string, bool) {
	timeNow := time.Now()
	userBirthdayDateMonth, userBirthdayDateDay := userBirthdayDate.Month(), userBirthdayDate.Day()
	yearNow, monthNow, dayNow := timeNow.Year(), timeNow.Month(), timeNow.Day()
	dateNow := time.Date(yearNow, monthNow, dayNow, 0, 0, 0, 0, time.UTC)
	userBirthdayDateThisYear := time.Date(yearNow, userBirthdayDateMonth, userBirthdayDateDay, 0, 0, 0, 0, time.UTC)

	switch days := userBirthdayDateThisYear.Sub(dateNow).Hours() / 24; {
	case days < 0:
		userDateOfBirthNextYear := userBirthdayDateThisYear.AddDate(1, 0, 0)
		days := userDateOfBirthNextYear.Sub(dateNow).Hours() / 24
		return fmt.Sprintf("Hello, %s! Your birthday is in %v day(s)", username, days), true
	case days == 0:
		return fmt.Sprintf("Hello, %s! Happy birhday!", username), true
	case days > 0:
		return fmt.Sprintf("Hello, %s! Your birthday is in %v day(s)", username, days), true
	default:
		return fmt.Sprintln("Something went wrong with birthday calculation"), false
	}
}
