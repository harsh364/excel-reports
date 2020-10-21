Excel report Generator helps creating excel reports from biq query with minimal efforts and development required

# Table of Contents

- [Overview](#overview)
- [Installing](#installing)
- [Getting Started](#getting-started)
  * [Report Structure and Code](#report-structure-and-code)
    * [Query Builder and Fetcher](#query-builder-and-fetcher)
    * [Schema](#schema)
    * [Utils](#utils)
    * [Columns and Sheets](#columns-and-sheets)
    * [Generator](#generator)
    * [API End Point](#api-end-point)
  * [Using the Excel Report Generator](#using-the-excel-report-generator)
  * [Commands](#commands)
- [Adding New Report](#adding-new-report)
- [To Do](#todo)

# Overview
Excel Report Generator is a tool that will help creating excel reports with minimal efforts and coding required

Excel Report Generator provides:
* Easy subcommand-based CLI for report generation
* Parallel report generation
* Easy development frame work for adding new report

# Installing
- Download and install golang from [Here](https://golang.org/doc/install#download)

Using Excel Report Generator is easy. Clone the repo, use `go install` to install executable
along with the library and its dependencies:

# Getting Started

While you are welcome to provide your own organization, typically an Excel Report
 will follow the following organizational structure:

```
  ▾ appName/
    ▾ cmd/
        main.go
    ▾ reports/
      ▾ your-report/
        ▾ bq/
           querybuilder.go
           fetchdata.go
        ▾ excel/
          ▾ sheets/
              sheet1.go
              sheet2.go
          ▾ formatters/
              sheet1.go
              sheet2.go
          ▾ columns/
              column1.go
              column2.go
        ▸ schema/
        ▸ utils
        generator.go
      ▾ shared/
        ▸ bq/
        ▸ excel/
        ▸ schema/
        ▸ utils
     ▸ utils
     api_endpoint_your_report.go
       
```
Typically main.go file is very bare. It serves of initilizing excel report generator and handles CLI logic
```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// args := os.Args[1:]
	fmt.Println()
	// argParser(args)
	run := flag.String("run", "", "Runs the provided report")
	list := flag.Bool("list", false, "Lists all the Reports")

	flag.Parse()
	if *list {
		fmt.Println(getList())
	}
	runReport(*run)
}

```
## Report Structure and Code

Each report can have a bq directory for queries and fetchers, an excel directory for all the sheets, columns and formatters, schema for handling report schema and utils for report specific utils. Any shared logic between multiple reports will go inside shared directory following same structure as of report.

### Query Builder and Fetcher

A Query builder generates bq query to fetch data, here is a sample querybuilder.go file

```go
package querybuilder

import (
	"example.com/emailreports/reports/shared/schema"
)

// SampleQuery returns sample query to fetch name and age data
func SampleQuery(qp schema.QueryParameter) string {
	return `
		-- Sample Query --
		SELECT 
			name, age
		FROM
			` + "`project_name.data_set_name.sample_table`" + `
		WHERE
			age > 20
		`
}

```

The query logic can be split into select, from, where and groupby. If the query is used or intended to be used accross multiple reports then the query builder should be placed in shared/bq/querbuilder/

Fetcher is used to fetch data from bq, while you are free to write your own fetcher, here is sample fetcher already implemented to fetch all types of data

```go
package fetchdata

import (
	"fmt"

	"cloud.google.com/go/bigquery"
	sharedSchema "example.com/emailreports/reports/shared/schema"
	"example.com/emailreports/utils"
	"google.golang.org/api/iterator"
)

// GetEntries fetches data from bq
func GetEntries(res chan sharedSchema.Entries, query string) {
	var bks sharedSchema.Entries
	it := utils.BqQuery(query)
	for {
		// var bk sharedSchema.Entry
		// err := it.Next(&bk)
		var m map[string]bigquery.Value
		err := it.Next(&m)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("Error in parsing: Fetched Data: ", err)
		}
		bks = append(bks, sharedSchema.GetEntry(m))
	}
	res <- bks
}

```

The fetcher is already implemented for general purpose, you just have to add your entries in shared/schema/entry.go file.

### Schema

Any custom schema/type defined can be placed in schema directory will all the methods defined in the same file. Below is the sheet schema which can be used accross multiple reports and is hence placed in shared/schema/sheet.go

```go
package schema

import (
	"reflect"

	sharedUtils "example.com/emailreports/reports/shared/utils"
	"example.com/emailreports/utils"
)

// Sheet is a struct which contains
// Name of the sheet
// Data which are grouping rows
// Headers which are columns
type Sheet struct {
	Name         sharedUtils.Sheet
	Data         GroupingRows
	SuperHeaders []utils.SuperHeader
	SubHeaders   SubHeaders
	Info         string
}

// Sheets is an array of sheet
type Sheets []Sheet

// GenerateSheet returns a sheet
func GenerateSheet(name sharedUtils.Sheet, data GroupingRows, superHeaders []utils.SuperHeader, subHeaders SubHeaders, info string) Sheet {
	return Sheet{
		Name:         name,
		Data:         data,
		SuperHeaders: superHeaders,
		SubHeaders:   subHeaders,
		Info:         info,
	}
}

// Equals compares two sheets
func (sheet Sheet) Equals(sheet2 Sheet) bool {
	// if &sheet == &sheet2 {
	// 	return true
	// }
	if sheet.Name != sheet2.Name ||
		!reflect.DeepEqual(sheet.SubHeaders, sheet2.SubHeaders) ||
		!reflect.DeepEqual(sheet.SuperHeaders, sheet2.SuperHeaders) ||
		!sheet.Data.Equal(sheet2.Data) {
		return false
	}
	return true
}

// ConvertToTwoDecimal the sheet numbers to two decimal
func (sheet Sheet) ConvertToTwoDecimal() Sheet {
	return GenerateSheet(sheet.Name, sheet.Data.ConvertToTwoDecimal(), sheet.SuperHeaders, sheet.SubHeaders, sheet.Info)
}

```

### Utils

All the general utility functions and constants can be directly placed inside utils directory available at the project root, while utils and constants specific to single report will be placed in reports/your-report/utils/ whereas any constant or utility function common accross the reports can be placed in reports/shared/utils/ Here are some sample utility functions

```go
package utils

import (
	"fmt"
	"strings"
	"time"
)

// StrListToStr joins list of string using comma required for BigQuery
// Input => ["abCd", "DEfg"]
// Output  => " \"abcd\", \"defg\" "
func StrListToStr(arr []string) string {
	return strings.ToLower("\"" + strings.Join(arr[:], "\",\"") + "\"")
}

// IntListToStr joins list of int using comma required for BigQuery
// Input => [1, 2, 3]
// Output  => "1, 2, 3"
func IntListToStr(intarr interface{}) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(intarr)), ","), "[]")
}

// MnthYr return map of year and months array from startDate and endDate
// Input => startDate = "2019-11-01", endDate = "2020-02-01"
// Output => { 2019: [11, 12], 2020: [1,2] }
func MnthYr(startDate, endDate time.Time) map[int][]int {
	var mnthYr = make(map[int][]int)
	endYr, endTimeMth, _ := endDate.Date()
	stYr, stTimeMth, _ := startDate.Date()
	stMth := int(stTimeMth)
	endMth := int(endTimeMth)
	for stYr < endYr || (stYr == endYr && stMth <= endMth) {
		mnthArr := mnthYr[stYr]
		if mnthArr == nil {
			mnthArr = []int{stMth}
		} else {
			mnthArr = append(mnthArr, stMth)
		}
		mnthYr[stYr] = mnthArr
		stMth++
		if stMth > 12 {
			stMth = 1
			stYr++
		}
	}
	return mnthYr
}

```

### Columns and Sheets

Column directory contains logic to fetch specific column and map the response in genric response structure, here is a sample column file

```go
package columns

import (
	"example.com/emailreports/app"
	querybuilder "example.com/emailreports/reports/shared/bq"
	sharedSchema "example.com/emailreports/reports/shared/schema"
	sharedUtils "example.com/emailreports/reports/shared/utils"
)

// Weight returns average weight
func Weight(qp sharedSchema.QueryParameter) sharedSchema.Entries {
	var col sharedUtils.Column

	res := qp.Fetcher.GetEntries(querybuilder.GetQuery(qp, qp.Selections, app.Configs.Table))

	for i, re := range res {
		re.Column = col
		res[i] = re
	}
	return res
}

```

Each sheet is made up of multiple columns with there aggregation and/or average or total row added, sheet files contain complete aggregation and generation logic for columns within the sheets along with sheet headers and sub headers. 

*Column Fetcher* : used to fetch all the columns below is a sample column fetcher
```go
// RaceWiseColumnFetchers fetches columns for Race Wise sheet
var RaceWiseColumnFetchers = sharedSchema.ColumnFetcher{
	sharedUtils.AvgWeightolumn: sharedSchema.EntriesFetch{},
}
```

Every sheet will have its column data collector defined with execute method written on it, below is sample code
```go
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
```
Each sheet file will have an addToRow function used to convert the entries into excel rows and group them according to the requirement.
```go
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
	grpRows.Grouping(
		res,
		attr,
		[]sharedUtils.ColumnRowValue{
			{Column: sharedUtils.NoneColumn, RowValue: "", Row: sharedSchema.Row{sharedUtils.RaceColumn: sharedUtils.Average}},
			{Column: sharedUtils.NoneColumn},
		},
		sharedSchema.Row{},
		"transpose2",
		sharedSchema.GetDivideBy(),
	)
}
```

Every sheet will have its super header and sub headers, these can either be shared among the sheets or specific to a sheet, below is an example

```go
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
```
Finally all these methods will be used to generate a sheet in the Get function for your sheet.
Below is a combined sample sheet code

```go
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
	grpRows.Grouping(
		res,
		attr,
		[]sharedUtils.ColumnRowValue{
			{Column: sharedUtils.NoneColumn, RowValue: "", Row: sharedSchema.Row{sharedUtils.RaceColumn: sharedUtils.Average}},
			{Column: sharedUtils.NoneColumn},
		},
		sharedSchema.Row{},
		"transpose2",
		sharedSchema.GetDivideBy(),
	)
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

```

### Generator

The generator is used to generate excel report, it generates each sheet parallely and formats the sheets, the generator contains sheet fetchers and data collectors which are similar to the column data collectors and fetchers, here is a sample generator code

```go
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
	case sharedUtils.SummaryRevenueDetailed:
		go sheets.GetSummaryRevenueDetailedSheet(bqDatesMap, mnth, yr, date, respChan, totalRowName, fetcher, filters)
	}
	return <-respChan
}

// SheetFetchers Demo report Sheet column fetchers
var SheetFetchers = map[sharedUtils.Sheet]sharedSchema.ColumnFetcher{
	sharedUtils.RaceWiseWeight:         sheets.RaceWiseColumnFetchers,
	sharedUtils.SummaryRevenueDetailed: sheets.SummaryRevenueDetailedColumnFetchers,
}

// GetDemoReport returns sheets
func GetDemoReport(date time.Time, reportName string, sheetFetchers map[sharedUtils.Sheet]sharedSchema.ColumnFetcher, filters sharedSchema.FilterClauses, totalRowName string) {
	sheets := []sheetGenerator{
		{
			sheetName: sharedUtils.RaceWiseWeight, date: date, fetcher: sheetFetchers[sharedUtils.RaceWiseWeight], filters: filters, totalRowName: totalRowName,
		},
		{
			sheetName: sharedUtils.SummaryRevenueDetailed, bqDatesMap: bqDatesMap, date: date, mnth: mnth, yr: yr, wks: wks, fetcher: sheetFetchers[sharedUtils.SummaryRevenueDetailed], filters: filters,
		},
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
		case sharedUtils.SummaryRevenueDetailed:
			formatters.SummaryRevenueDetailedSheet(f, sheet, totalRowName, asOfDisplayDate, "reportName")
		}
	}

	utils.SetActiveSheetAndDeleteDefaultSheet(f, string(sharedUtils.RaceWiseWeight))

	if err := f.SaveAs(filePath); err != nil {
		println(err.Error())
	}
}

```

### API End Point

Each report will have it's api endpoint where it will get features like report name, filters, report generation date etc. Below is sample api end point code

```go
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

```

## Using the Excel Report Generator

Using the Excel report generator is fairly easy and requires you to run the program with command line arguments.
Use the following command to build and run the program and generate specific report:

`go run ./cmd run "report name"`

## Commands

Currently there are only 3 commands
* *run*   : Used to generate specific report
* *list*  : Lists all the reports
* *help*  : Description of all the commands with usage
``

# Adding New Report

### Adding new report is divided into 5 parts
1. Building query
2. Fetching data
3. Grouping columns
4. Adding Columns to sheet
5. Creating End point and generator

### Building Query
1. Add your query type and selections in report/shared/schema/query_parameter
2. Add your query type case and selections cases in reports/shared/bq/querybuilder

### Fetching Data
1. Prepare your query parameters by adding query type, selections, filters etc in reports/your-report/excel/sheet/your-sheet.go
1. Add your column case with data type and attribute in reports/shared/schema/entry
2. Update all the functions/Getters for column,entry and stribute in reports/shared/schema/entry
3. Add your column to reports/shared/utils/columns
4. Create column fetcher function file in reports/your-report/excel/columns/
5. Write column fetcher, column data collector and execute method for the data collector

### Grouping Columns
1. Create a function similar to addToRows function in reports/demoreport/excel/sheets/race_wise_weight.go
2. Add Columns which will be used to group
3. Use the function to create row groups for excel sheet

### Adding Columns to sheet
1. Create Super headers and Sub Headers for the sheet
2. Write formatter for your sheet, defining start column, end column header style etc in reports/demoreport/excel/formatters/your-sheet.go
3. You can add or change color, width, header style etc by adding your style in utils/excel_helper.go

### Creating End Point and Generator
1. Create Sheet Fetchers, Sheet generator and execute method for generator
2. Add all the sheets in cases for execute and formatter
3. Create generator call from the api endpoint by creating an endpoint file
4. Add your report in the list of reports in cmd/main.go

# Todo
