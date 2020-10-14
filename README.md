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
1. Add your query type and selections in shared/schema/query_parameter
2. Add your query type case and selections cases in shared/bq/querybuilder
