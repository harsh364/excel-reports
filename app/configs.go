package app

import (
	"log"
	"os"
)

// dataTable is a table for testing
const dataTable = "`bigquery-public-data.samples.natality`"

// Configs is struct for app configurations
var Configs struct {
	ENV, Table, Project, Dataset string
}

func init() {
	log.Println("Config Initializing... ")
	Setup()
	log.Printf("%+v \n", Configs)
	log.Println("Initialization Completed ")
}

// Setup sets up the configs
func Setup() {
	Configs.ENV = getValue(os.Getenv("ENV"), "DEV")
	Configs.Table = getValue(os.Getenv("TABLE"), dataTable)
	Configs.Project = getValue(os.Getenv("PROJECT"))
	Configs.Dataset = getValue(os.Getenv("DATASET"), "sample")
}

func getValue(opts ...string) string {
	var val string
	for _, v := range opts {
		if len(v) > 0 {
			val = v
			break
		}
	}
	return val
}
