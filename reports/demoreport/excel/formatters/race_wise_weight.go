package formatters

import (
	"example.com/emailreports/reports/shared/schema"
	sharedUtils "example.com/emailreports/reports/shared/utils"
	"example.com/emailreports/utils"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// RaceWiseWeightSheet creates monthly formated work sheet and attach to the given excel file
func RaceWiseWeightSheet(f *excelize.File, sheet schema.Sheet, totalRowName, asOfDisplayDate, reportName string) (int, string) {
	rowNum := 2
	startCol := "B"
	headerEndCol := "C"
	sheetName := string(sheet.Name)
	f.NewSheet(sheetName)

	superHeaders := sheet.SuperHeaders
	utils.AddSuperHeaderToSheet(superHeaders, f, sheetName, &rowNum, utils.NoColorHeaderSty)
	utils.SetRow(f, sheetName, startCol, &rowNum, sheet.SubHeaders.Headers(), utils.Report4HeaderSty)

	totalGrpRow, grpRows := sheet.Data.ExtractTotalRow("")
	for _, r := range grpRows.Sort() {
		setGroupingRow(f, sheetName, startCol, headerEndCol, r, &rowNum, sheet.SubHeaders.Columns())
	}
	// rowNum++
	setGroupingRow(f, sheetName, startCol, headerEndCol, totalGrpRow, &rowNum, sheet.SubHeaders.Columns())
	return rowNum, "B"
}

func setGroupingRow(f *excelize.File, sheetName, startCol, headerEndCol string, r schema.GroupingRow, rowNum *int, headers sharedUtils.Columns) {
	// Add all Children Row
	index := 1
	for _, c := range r.Children.Sort() {
		updatedc := c.ToArray(headers)
		if index > 1 {
			updatedc[0] = ""
		}
		utils.SetRow(f, sheetName, startCol, rowNum, updatedc, utils.RoundOffToZeroSty)
		index++
	}
	// Add Parent Row
	utils.SetRow(f, sheetName, startCol, rowNum, r.Parent.ToArray(headers), utils.RoundOffToZeroTotalSty)
}
