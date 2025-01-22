package main

import (
	"fmt"
	"os"
	"time"
)

// validates if the deadline-string can be parsed to a valid date
func validateDeadline(deadline string) (time.Time, error) {
	parsedDeadline, err := time.Parse("2006-01-02", deadline)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid deadline format, expected yyyy-mm-dd")
	}
	return parsedDeadline, nil
}

// calculates the difference between the deadline and now
func calculateDifference(deadline time.Time) time.Duration {
	today := time.Now()
	return deadline.Sub(today)
}

func main() {
	// is the argument is provided correctly, if not return
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./main <deadline> ")
		fmt.Println("Example: ./main 2025-01-10 (Format yyyy-mm-dd)")
		os.Exit(1)
	}

	// Hacky way to get the deadline from the argument
	deadlineStr := os.Args[1]

	// check deadline and exits if there's an error
	deadline, err := validateDeadline(deadlineStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// is deadline is in the past and exits if it is one
	if deadline.Before(time.Now()) {
		fmt.Println("The deadline should be in the future.")
		os.Exit(1)
	}

	// difference in days
	diff := calculateDifference(deadline)
	fmt.Printf("You have %.0f days before your deadline.\n", diff.Hours()/24)
}
