package utils

import "time"

func RetrieveCurrentDate(date string) (time.Time, error) {
	// Define the layout according to the given date format
	layout := "2006-01-02 15:04:05"
	// Parse the date string using the specified layout
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}
	return parsedDate, nil
}
