package sheets

import (
	"time"

	"example.com/emailreports/reports/demoreport/excel/columns"
	sharedSchema "example.com/emailreports/reports/shared/schema"
	sharedUtils "example.com/emailreports/reports/shared/utils"
	"example.com/emailreports/utils"
	"github.com/JigneshSatam/parallel"
)

// RaceWiseColumnFetchers current month column fetchers
var RaceWiseColumnFetchers = sharedSchema.ColumnFetcher{
	sharedUtils.AvgWeightolumn: sharedSchema.EntriesFetch{},
}

type colDataCollector struct {
	col sharedUtils.Column
	qp  sharedSchema.QueryParameter
}

type executeOutput struct {
	result sharedSchema.Entries
	schema sharedSchema.EntryAttribute
}

func (e colDataCollector) Execute() interface{} {
	result := sharedSchema.Entries{}
	response := executeOutput{}
	schema := sharedSchema.Weight
	qp := e.qp
	switch e.col {
	case sharedUtils.AvgWeightolumn:
		result = columns.Weight(qp)
		response = executeOutput{result: result, schema: schema}
	}
	return response
}

// GetRaceWiseWeight returns race by weight of child sheet
func GetRaceWiseWeight(date time.Time, sheetChan chan sharedSchema.Sheet, totalRowName string, sheetColumns sharedSchema.ColumnFetcher, filters sharedSchema.FilterClauses) {
	qp := sharedSchema.QueryParameter{
		ReportDate:    date,
		FilterClauses: filters,
		Type:          sharedSchema.WeightQueryType,
		Selections: sharedSchema.Selections{
			sharedSchema.RaceSelection,
			sharedSchema.FatherAgeSelection,
		},
	}

	defer close(sheetChan)

	grpRows := make(sharedSchema.GroupingRows)
	for resArr := range rwwColumns(qp, sheetColumns) {
		schema := resArr.(executeOutput).schema
		res := resArr.(executeOutput).result
		// fmt.Printf("%#v", res)
		if len(res) > 0 {
			addToRows(grpRows, schema, res, totalRowName)
		}
	}

	sheetChan <- sharedSchema.GenerateSheet(
		sharedUtils.RaceWiseWeight,
		grpRows,
		rwwSuperHeaders(),
		rwwSubHeaders(),
		string(sharedUtils.RaceWiseWeight),
	)
}

func rwwColumns(qp sharedSchema.QueryParameter, sheetColumns sharedSchema.ColumnFetcher) <-chan interface{} {
	executors := []colDataCollector{}
	for col, fetcher := range sheetColumns {
		qp.Fetcher = fetcher
		executors = append(executors, colDataCollector{col, qp})
	}
	return parallel.Run(executors)
}

func addToRows(grpRows sharedSchema.GroupingRows, attr sharedSchema.EntryAttribute, res sharedSchema.Entries, totalRowName string) {
	grpRows.Grouping(
		res,
		attr,
		[]sharedUtils.ColumnRowValue{
			{Column: sharedUtils.RaceColumn, RowValue: totalRowName},
			{Column: sharedUtils.FatherAgeColumn},
		},
		sharedSchema.Row{},
		"transpose2",
		sharedSchema.GetDivideBy(),
	)
	// grpRows.Grouping(
	// 	res,
	// 	attr,
	// 	[]sharedUtils.ColumnRowValue{
	// 		{Column: sharedUtils.NoneColumn, RowValue: "", Row: sharedSchema.Row{sharedUtils.RaceColumn: sharedUtils.Average}},
	// 		{Column: sharedUtils.NoneColumn},
	// 	},
	// 	sharedSchema.Row{},
	// 	"transpose2",
	// 	sharedSchema.GetDivideBy(),
	// )
}

func rwwSuperHeaders() []utils.SuperHeader {
	arr := []utils.SuperHeader{}
	return arr
}

func rwwSubHeaders() sharedSchema.SubHeaders {

	cols := sharedSchema.SubHeaders{
		{Column: sharedUtils.RaceColumn, Header: "Race", Group: sharedUtils.Group1},
		{Column: sharedUtils.FatherAgeColumn, Header: "Father's Age", Group: sharedUtils.Group1},
		{Column: sharedUtils.AvgWeightolumn, Header: "Average Weight", Group: sharedUtils.Group1},
	}
	return cols
}
