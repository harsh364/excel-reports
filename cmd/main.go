package main

import (
	"fmt"
	"time"

	excelreports "example.com/emailreports"
)

func main() {
	today := time.Now()
	fmt.Println("Today: ", today)
	// date := time.Date(2020, time.July, 16, 0, 0, 0, 0, time.UTC)
	// date := time.Now()
	// Display(date)
	excelreports.GenDemoReport()

	fmt.Println("Report Generated.  Time taken: ", time.Since(today))
}
