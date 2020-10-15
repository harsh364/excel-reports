# excel-reports
Excel report helps creating excel reports from biq query with minimal efforts and development required

## Adding new report

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
