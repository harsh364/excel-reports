package querybuilder

import (
	"strings"

	"example.com/emailreports/app"
	"example.com/emailreports/reports/shared/schema"
	"example.com/emailreports/utils"
)

var columnMapping = map[string]string{
	"child_race":    "child_race",
	"weight_pounds": "weight_pounds",
	"father_age":    "father_age",
	"mother_age":    "mother_age",
}

// FilterWeightQuery returns Weight using provided filters
func FilterWeightQuery(qp schema.QueryParameter) string {
	selectClause, groupingClause := selectionGroupingClause(qp, qp.Selections)
	return `-- ` + qp.Name + ` --
		` + selectClause + `
		` + groupingClause
}

// =================================================
// ========= Selection and Grouping Clause =========
// =================================================

func selectionGroupingClause(qp schema.QueryParameter, selections schema.Selections) (string, string) {
	selectClause := "SELECT\n"
	grouping := "GROUP BY\n"

	selectClause += " avg(weight_pounds) as weight_pounds"

	for _, sel := range selections {
		switch sel {
		case schema.RaceSelection:
			selectClause += ", child_race AS child_race\n"
			grouping += "child_race,"
		case schema.FatherAgeSelection:
			selectClause += ", father_age AS father_age\n"
			grouping += "father_age,"
		case schema.MotherAgeSelection:
			selectClause += ", mother_age AS mother_age\n"
			grouping += "mother_age,"
		}
	}

	selectClause += ` FROM` + app.Configs.Table
	return selectClause, strings.TrimSuffix(grouping, ",")
}

// =================================================
// ==================== Filters ====================
// =================================================

func filterClauses(prefix string, fcs schema.FilterClauses) string {
	filters := ""
	for _, f := range fcs {
		if columnName, present := columnMapping[f.Attribute]; present {
			filters += `
			` + prefix + `
				(
					LOWER(` + columnName + ")" + operator(f.Inclusion) + "(" + utils.IntListToStr(f.Options) + ")" + `
					` + nullClause(f.NullRequired, columnName) + `
				)`
			prefix = "AND"
		}
	}
	return filters
}

// =================================================
// ================ Helper functions ===============
// =================================================

func operator(inclusion bool) string {
	op := " IN "
	if !inclusion {
		op = " NOT IN "
	}
	return op
}

func nullClause(reqired bool, columnName string) string {
	if reqired {
		return " OR " + columnName + " IS NULL"
	}
	return ""
}
