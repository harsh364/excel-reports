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
