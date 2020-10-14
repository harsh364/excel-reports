package utils

import (
	"context"
	"fmt"

	"example.com/emailreports/app"

	"cloud.google.com/go/bigquery"
)

// ProjectID is the BigQuery project ID
var ProjectID = app.Configs.Project

// DatasetName is the BigQuery DataBase name
var DatasetName = app.Configs.Dataset
var client *bigquery.Client
var bqDataset *bigquery.Dataset
var ctx context.Context

func init() {
	var err error
	ctx = context.Background()
	client, err = bigquery.NewClient(ctx, ProjectID)
	if err != nil {
		fmt.Println("Error in init: ", err.Error())
	}
	bqDataset = client.DatasetInProject(ProjectID, DatasetName)
}

// BqQuery fetches the data from BigQuery
func BqQuery(q string) *bigquery.RowIterator {
	query := client.Query(q)
	it, err := query.Read(ctx)
	if err != nil {
		fmt.Println("Error in reading: ", err)
	}
	return it
}

// BQDryRun check if the query is valid
func BQDryRun(q string) error {
	query := client.Query(q)
	query.DryRun = true
	job, err := query.Run(ctx)
	if err != nil {
		return err
	}
	// Dry run is not asynchronous, so get the latest status and statistics.
	status := job.LastStatus()
	if err != nil {
		return err
	}
	fmt.Printf("This query will process %d bytes\n", status.Statistics.TotalBytesProcessed)
	return nil
}

// BqInsert inserts the data to BigQuery
func BqInsert(tableName string, src interface{}) {
	table := bqDataset.Table(tableName)
	ins := table.Inserter()
	err := ins.Put(ctx, src)
	if err != nil {
		fmt.Println("Error in creating user resport status: ", err)
	} else {
		fmt.Println("Recorder Created.")
	}
}
