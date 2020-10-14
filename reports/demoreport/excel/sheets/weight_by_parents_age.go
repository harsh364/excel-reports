package sheets

// import (
// 	"strconv"
// 	"time"

// 	"example.com/emailreports/reports/demoreport/schema"
// 	sharedColumns "example.com/emailreports/reports/shared/excel/columns"
// 	sharedSchema "example.com/emailreports/reports/shared/schema"
// 	sharedUtils "example.com/emailreports/reports/shared/utils"
// 	"example.com/emailreports/utils"
// 	"example.com/emailreports/utils/disneycalendar"
// 	"github.com/JigneshSatam/parallel"
// )

// // SummaryColumnFetchers current month column fetchers
// var SummaryColumnFetchers = sharedSchema.ColumnFetcher{
// 	sharedUtils.Weight:                sharedSchema.RevenueEntriesFetch{},
// }

// type colDataCollector struct {
// 	col        sharedUtils.Column
// 	qp         sharedSchema.QueryParameter
// }

// type executeOutput struct {
// 	result []sharedSchema.RevenueEntries
// 	schema sharedSchema.RevenueEntryAttribute
// }

// func (e colDataCollector) Execute() interface{} {
// 	result := []sharedSchema.RevenueEntries{}
// 	response := executeOutput{}
// 	qp := e.qp
// 	switch e.col {
// 	case sharedUtils.Weight:
// 		result = columns.Weight(qp)
// 		response = executeOutput{result: result, schema: sharedSchema.Revenue}
// 	}
// 	return response
// }

// // GetRaceWiseWeight returns race by weight of child sheet
// func GetRaceWiseWeight(bqDatesMap disneycalendar.BQDatesMap, mnth utils.Month, yr utils.Year, date time.Time, wks []int, sheetChan chan sharedSchema.Sheet, totalRowName string, sheetColumns sharedSchema.ColumnFetcher, filters sharedSchema.FilterClauses) {
// 	qp := sharedSchema.QueryParameter{
// 		ReportDate:    date,
// 		FilterClauses: filters,
// 		Selections: sharedSchema.Selections{
// 			sharedSchema.RaceSelection,
// 		},
// 	}

// 	defer close(sheetChan)

// 	grpRows := make(sharedSchema.GroupingRows)
// 	for resArr := range summaryColumns(bqDatesMap, mnth, yr, qp, sheetColumns, wks) {
// 		schema := resArr.(executeOutput).schema
// 		for _, res := range resArr.(executeOutput).result {
// 			if len(res) > 0 {
// 				addToRows(revGrpRows, res, schema, totalRowName)
// 			}
// 		}
// 	}

// 	sheetChan <- sharedSchema.GenerateSheet(
// 		sharedUtils.RaceWiseWeight,
// 		revGrpRows,
// 		rwwSuperHeaders(len(wks)),
// 		rwwSubHeaders(wks, int(mnth), int(yr)),
// 		string(sharedUtils.RaceWiseWeight),
// 	)
// }

// func rwwColumns(bqDatesMap disneycalendar.BQDatesMap, mnth utils.Month, yr utils.Year, qp sharedSchema.QueryParameter, sheetColumns sharedSchema.ColumnFetcher, wks []int) <-chan interface{} {
// 	executors := []colDataCollector{}
// 	for col, fetcher := range sheetColumns {
// 		qp.Fetcher = fetcher
// 		executors = append(executors, colDataCollector{col, bqDatesMap, mnth, yr, qp, wks})
// 	}
// 	return parallel.Run(executors)
// }

// func addToRows(grpRows sharedSchema.GroupingRows, res sharedSchema.RevenueEntries, attr sharedSchema.RevenueEntryAttribute, totalRowName string) {
// 	divideBy := sharedSchema.GetDivideBy()
// 	grpRows.Grouping(
// 		res,
// 		attr,
// 		[]sharedUtils.ColumnRowValue{
// 			{Column: sharedUtils.ChannelName, RowValue: totalRowName},
// 			{Column: sharedUtils.RegionColumn},
// 		},
// 		sharedSchema.Row{},
// 		"transpose2",
// 		divideBy,
// 	)
// 	grpRows.Grouping(
// 		res,
// 		attr,
// 		[]sharedUtils.ColumnRowValue{
// 			{Column: sharedUtils.NoneColumn, RowValue: "", Row: sharedSchema.Row{sharedUtils.RegionColumn: sharedUtils.Total}},
// 			{Column: sharedUtils.NoneColumn},
// 		},
// 		sharedSchema.Row{},
// 		"transpose2",
// 		divideBy,
// 	)
// }

// func rwwSuperHeaders(numberOfWeeks int) []utils.SuperHeader {
// 	totalColum := 1
// 	pos := 8
// 	arr := []utils.SuperHeader{
// 		{Name: "Revenue", Position: utils.GetPosition(pos), RepeatFor: utils.PositionIncr(&pos, (numberOfWeeks + totalColum))},
// 		{Name: "ASR", Position: utils.GetPosition(pos), RepeatFor: utils.PositionIncr(&pos, (numberOfWeeks + totalColum))},
// 	}
// 	return arr
// }

// func rwwSubHeaders(wks []int, mnth, yr int) sharedSchema.SubHeaders {

// 	cols := sharedSchema.SubHeaders{
// 		{Column: sharedUtils.ChannelName, Header: "Race", Group: sharedUtils.Group1},
// 		{Column: sharedUtils.RegionColumn, Header: "Average Weight", Group: sharedUtils.Group1},
// 	}
// 	return cols
// }
