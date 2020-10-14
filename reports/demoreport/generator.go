package demoreport

import (
	"fmt"
	"os"
	"time"

	"example.com/emailreports/reports/demoreport/excel/formatters"
	"example.com/emailreports/reports/demoreport/excel/sheets"
	sharedSchema "example.com/emailreports/reports/shared/schema"
	sharedUtils "example.com/emailreports/reports/shared/utils"
	"example.com/emailreports/utils"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/JigneshSatam/parallel"
)

type sheetGenerator struct {
	sheetName    sharedUtils.Sheet
	totalRowName string
	date         time.Time
	fetcher      sharedSchema.ColumnFetcher
	filters      sharedSchema.FilterClauses
}

func (sheetGen sheetGenerator) Execute() interface{} {
	respChan := make(chan sharedSchema.Sheet)
	date := sheetGen.date
	totalRowName := sheetGen.totalRowName
	fetcher := sheetGen.fetcher
	filters := sheetGen.filters
	switch sheetGen.sheetName {
	case sharedUtils.RaceWiseWeight:
		go sheets.GetRaceWiseWeight(date, respChan, totalRowName, fetcher, filters)
		// case sharedUtils.SummaryRevenueDetailed:
		// 	go sheets.GetSummaryRevenueDetailedSheet(bqDatesMap, mnth, yr, date, respChan, totalRowName, fetcher, filters)
	}
	return <-respChan
}

// SheetFetchers Demo report Sheet column fetchers
var SheetFetchers = map[sharedUtils.Sheet]sharedSchema.ColumnFetcher{
	sharedUtils.RaceWiseWeight: sheets.RaceWiseColumnFetchers,
	// sharedUtils.SummaryRevenueDetailed: sheets.SummaryRevenueDetailedColumnFetchers,
}

// GetDemoReport returns sheets
func GetDemoReport(date time.Time, reportName string, sheetFetchers map[sharedUtils.Sheet]sharedSchema.ColumnFetcher, filters sharedSchema.FilterClauses, totalRowName string) {
	sheets := []sheetGenerator{
		{
			sheetName: sharedUtils.RaceWiseWeight, date: date, fetcher: sheetFetchers[sharedUtils.RaceWiseWeight], filters: filters, totalRowName: totalRowName,
		},
		// {
		// 	sheetName: sharedUtils.SummaryRevenueDetailed, bqDatesMap: bqDatesMap, date: date, mnth: mnth, yr: yr, wks: wks, fetcher: sheetFetchers[sharedUtils.SummaryRevenueDetailed], filters: filters,
		// },
	}

	// currMntStDtCMD, currMntEnDtCMD := bqDatesMap.MonthStartEndDate(mnth, yr)
	asOfDisplayDate := date.Format(utils.AsOfDateLayout)

	fileName := "Demo Report-" + asOfDisplayDate + ".xlsx"
	filePath := "./" + fileName
	if os.Getenv("ENV") == "TEST" || os.Getenv("ENV") == "UAT" || os.Getenv("ENV") == "PROD" {
		fileName = reportName + " " + asOfDisplayDate + ".xlsx"
		filePath = "/tmp/" + fileName
	}

	f := excelize.NewFile()
	for sheet := range parallel.Run(sheets) {
		sheet := sheet.(sharedSchema.Sheet)
		switch sheet.Name {
		case sharedUtils.RaceWiseWeight:
			_, col := formatters.RaceWiseWeightSheet(f, sheet, totalRowName, asOfDisplayDate, "reportName")
			fmt.Println(col)
			// case sharedUtils.SummaryRevenueDetailed:
			// 	formatters.SummaryRevenueDetailedSheet(f, sheet, "totalRowName", asOfDisplayDate, "reportName")
		}
	}

	utils.SetActiveSheetAndDeleteDefaultSheet(f, string(sharedUtils.RaceWiseWeight))

	if err := f.SaveAs(filePath); err != nil {
		println(err.Error())
	}
}
