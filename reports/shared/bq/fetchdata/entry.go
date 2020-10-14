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
