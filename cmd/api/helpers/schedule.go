package helpers

import (
	"errors"
	"fmt"
	"time"
)

type Schedule uint8

const (
	Diario Schedule = iota
	Semanal
	Mensal
	Trimestral
	Semestral
	Anual
)

const (
	Day   time.Duration = 24 * time.Hour
)

func ParseCron(agenda Schedule, date time.Time) (string, error) {
	day := date.Day()
	week := int(date.Weekday())

	switch agenda {
	case Diario:

		return "0 0 * * *", nil
	case Semanal:
		return fmt.Sprintf("0 0 * * %d", week), nil
	case Mensal:
		return fmt.Sprintf("0 0 %d * *", day), nil
	case Trimestral:
		return fmt.Sprintf("0 0 %d 1-12/3 *", day), nil
	case Semestral:
		return fmt.Sprintf("0 0 %d 1-12/6 *", day), nil
	case Anual:
		return fmt.Sprintf("0 0 %d 1-12/12 *", day), nil
	}
	return "", errors.New("invalid schedule")
}

// GetFirstSunday returns the date of the first Sunday of the given year.
//
// It calculates the first day of the year and subtracts the number of days
// corresponding to its weekday to find the nearest preceding Sunday.
//
// Parameters:
//   - year: The year for which the first Sunday is determined.
//
// Returns:
//   - A time.Time object representing the first Sunday of the given year.
func GetFirstSunday(year int) time.Time {
	firstDay := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	week := firstDay.Weekday()
	if week <= 3 {
		return firstDay.Add(-Day * time.Duration(week))
	}
	return firstDay.Add(Day * time.Duration(7 - week))

	
}


// getWeek calculates the epidemiological week number and corresponding year
// based on Brazil's Ministry of Health epidemiological calendar.
//
// The first epidemiological week of the year is defined as the one containing
// at least four days of the new year, with each subsequent week having exactly
// seven days. This system is widely used for epidemiological surveillance.
//
// Parameters:
//   - date (time.Time): The date for which the epidemiological week is determined.
//
// Returns:
//   - year (int): The epidemiological year, which may differ from the calendar year.
//   - week (int): The epidemiological week number.
//
// Example:
//
//   date := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
//   year, week := getWeek(date)
//   fmt.Printf("Year: %d, Week: %d\n", year, week)
//
// This calculation is essential for consistent epidemiological data analysis.

func GetWeek(data time.Time) (year int, week int) {
	year, week = data.ISOWeek()
	if week == 52 {
		if 31-data.Day() >= 3 && data.Day() > 20 {
			week = 53
			return
		} else {
			week = 1
			year++
		}
	} else {
		year, week = data.Add(Day).ISOWeek()
	}
	return
}


func TimeIsBetween(t, min, max time.Time) bool {
    if min.After(max) {
        min, max = max, min
    }
    return (t.Equal(min) || t.After(min)) && (t.Equal(max) || t.Before(max))
}