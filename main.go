package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func parseDates() (since time.Time, till time.Time, error error) {
	if len(os.Args) < 3 {
		return since, till, errors.New("not enough arguments provided")
	}
	since, err := parse(1)
	if err != nil {
		return since, till, errors.New("error parsing the first date")
	}
	till, err = parse(2)
	if err != nil {
		return since, till, errors.New("error parsing the second date")
	}
	if !till.After(since) {
		return since, till, errors.New("second date has to be greater than the first")
	}
	return since, till, error
}

func parse(argIndex int) (timeTime time.Time, err error) {
	timeString := os.Args[argIndex]
	timeTime, err = time.Parse("2006-01-02", timeString)
	if err != nil {
		return timeTime, err
	}
	return timeTime, err
}

func days(d time.Duration) int {
	return int(d.Hours() / 24)
}

func main() {
	since, till, err := parseDates()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(-1)
	}
	difference := till.Sub(since)
	daysDifference := days(difference)
	fmt.Println(daysDifference)
}
