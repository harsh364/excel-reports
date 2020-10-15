Excel report Generator helps creating excel reports from biq query with minimal efforts and development required

# Table of Contents

- [Overview](#overview)
- [Installing](#installing)
- [Getting Started](#getting-started)
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
	"fmt"
	"os"
	"time"

	"example.com/emailreports/reports/shared/utils"

	excelreports "example.com/emailreports"
)

func main() {
	args := os.Args[1:]
	fmt.Println()
	argParser(args)
}

func argParser(args []string) {
	switch args[0] {
	case "-list":
		fmt.Println(getList())
	case "-help":
		fmt.Print(getHelp())
	case "-run":
		if len(args) > 1 {
			runReport(args[1])
		} else {
			fmt.Println("No report name provided to run, please provide a report name or use h for help")
		}
	default:
		fmt.Println("No argument Provided, please use -h for help on commands")
	}
}
```
File wise detailed documentation to be added

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
