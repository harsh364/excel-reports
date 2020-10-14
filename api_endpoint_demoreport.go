package excelreports

import (
	"time"

	"example.com/emailreports/reports/shared/utils"

	"example.com/emailreports/reports/demoreport"
	sharedSchema "example.com/emailreports/reports/shared/schema"
)

// GenDemoReport is function to generate demo report
func GenDemoReport() {
	// date := time.Date(2020, time.July, 01, 0, 0, 0, 0, time.UTC)
	date := time.Now()
	filters := sharedSchema.FilterClauses{
		{
			Attribute:    "child_race",
			Inclusion:    false,
			NullRequired: false,
			Options:      []string{"2"},
		},
	}
	reportName := "Demo Excel Report"
	demoreport.GetDemoReport(date, reportName, demoreport.SheetFetchers, filters, utils.Average)
}
