package formatters

import (
	"fmt"
	"strconv"

	"example.com/emailreports/reports/shared/schema"
	sharedUtils "example.com/emailreports/reports/shared/utils"
	"example.com/emailreports/utils"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// SetHeaderStyle sets header style
func SetHeaderStyle(f *excelize.File, sheetName, startCol, headerEndCol string, rowNum int) {
	f.SetCellStyle(sheetName, startCol+strconv.Itoa(rowNum), headerEndCol+strconv.Itoa(rowNum), utils.GetStyle(f, utils.HeaderSty))
}

// SetSheetStyle sets style to entire sheet as well as add extra columns
func SetSheetStyle(f *excelize.File, sheet, startCol string, subHdrs schema.SubHeaders) {
	grps := []sharedUtils.Group{sharedUtils.Group1, sharedUtils.Group2, sharedUtils.Group3}
	stCol, err := excelize.ColumnNameToNumber(startCol)
	if err != nil {
		fmt.Println("Error in ColumnNameToNumber", err)
	}
	enCol := stCol
	for _, grp := range grps {
		groupLen := len(subHdrs.FilterByGroup(grp))
		if groupLen == 0 {
			break
		}

		enCol = stCol + groupLen - 1
		endColName, err := excelize.ColumnNumberToName(enCol)
		startColName, err := excelize.ColumnNumberToName(stCol)
		f.SetColWidth(sheet, startColName, endColName, 15)

		blankColumn := enCol + 1
		blankColumnName, err := excelize.ColumnNumberToName(blankColumn)
		f.InsertCol(sheet, blankColumnName)
		if err != nil {
			fmt.Println("Error in ColumnNumberToExcelName", err)
		}
		f.SetColWidth(sheet, blankColumnName, blankColumnName, 2)
		stCol = blankColumn + 1
	}
}
