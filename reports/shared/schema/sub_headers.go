package schema

import (
	"strconv"
	"time"

	"example.com/emailreports/reports/shared/utils"
)

// SubHeader is a type for sub header
type SubHeader struct {
	Column utils.Column
	Header interface{}
	Group  utils.Group
}

// SubHeaders is an array of sub headers
type SubHeaders []SubHeader

// FilterByGroup filter SubHeaders by group
func (cols SubHeaders) FilterByGroup(grp utils.Group) SubHeaders {
	fixedCols := SubHeaders{}
	for _, col := range cols {
		if col.Group == grp {
			fixedCols = append(fixedCols, col)
		}
	}
	return fixedCols
}

// Headers return header array of the columns
func (cols SubHeaders) Headers() []interface{} {
	headers := []interface{}{}
	for _, col := range cols {
		headers = append(headers, col.Header)
	}
	return headers
}

// Columns return names array of the columns
func (cols SubHeaders) Columns() utils.Columns {
	hColumns := utils.Columns{}
	for _, col := range cols {
		hColumns = append(hColumns, col.Column)
	}
	return hColumns
}

// LastYrMnthStr returns last month-year sub header string
func LastYrMnthStr(mnth, yr int) string {
	return yrMnthStr(mnth, yr-1)
}

// CurrYrMnthStr returns current month-year sub header string
func CurrYrMnthStr(mnth, yr int) string {
	return yrMnthStr(mnth, yr)
}

func yrMnthStr(mnth, yr int) string {
	return YearStr(yr) + " " + time.Month(mnth).String()
}

// YearStr returns year super-header string
func YearStr(yr int) string {
	return "FY" + strconv.Itoa(yr)
}
