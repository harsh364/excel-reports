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
