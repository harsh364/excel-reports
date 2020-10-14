package utils

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// DefaultRowHeight is default height of the sheet rows
const DefaultRowHeight = 26.25

// SetRow add the row in the provided excel file i.e `f` into sheet i.e `sheet` at the given `row` pointer
// and increaments the row pointer value by one.
func SetRow(f *excelize.File, sheet, startCol string, row *int, vals []interface{}, sty style) {
	rowStr := strconv.Itoa(*row)
	startColNum, err := excelize.ColumnNameToNumber(startCol)
	if err != nil {
		fmt.Println("Error in ColumnNameToNumber", err)
	}
	endCol, err := excelize.ColumnNumberToName(startColNum + len(vals) - 1)
	if err != nil {
		fmt.Println("Error in ColumnNumberToExcelName", err)
	}
	f.SetSheetRow(sheet, startCol+rowStr, &vals)
	f.SetCellStyle(sheet, startCol+rowStr, endCol+rowStr, GetStyle(f, sty))
	f.SetRowHeight(sheet, *row, DefaultRowHeight)
	(*row)++
}

// SetColumnStyle is used to set a style for all the cells particular column from startRow to endRow for a
func SetColumnStyle(f *excelize.File, sheet, columnName string, startRow int, endRow int, sty style) {
	startRowStr := strconv.Itoa(startRow)
	endRowStr := strconv.Itoa(endRow)
	f.SetCellStyle(sheet, columnName+startRowStr, columnName+endRowStr, GetStyle(f, sty))
}

type style string

var styles = make(map[*excelize.File]map[style]int)

const (
	// HeaderSty style
	HeaderSty style = "header"
	// NoSty style
	NoSty style = ""
	//TotalSty  style
	TotalSty style = "total"
	// DefaultSty style
	DefaultSty style = "default"
	// InfoTextSty style
	InfoTextSty style = "infoText"
	// RoundOffToZeroSty style
	RoundOffToZeroSty style = "roundOffToZero"
	// RoundOffToZeroTotalSty style
	RoundOffToZeroTotalSty style = "roundOffToZeroTotal"
	// Report4HeaderSty style
	Report4HeaderSty style = "report4HeaderStyle"
	// NoColorHeaderSty style
	NoColorHeaderSty style = "noColorHeaderStyle"
	// DefaultGreyColor style
	DefaultGreyColor style = "defaultGreyColor"

	defaultStyleStr = `
		"alignment":{
			"horizontal":"center",
			"vertical":"center",
			"shrink_to_fit":false,
			"wrap_text":true
		}
	`
	roundOffToTwoStr = `
		"number_format": 2
	`
	commaNumberFormatStr = `
		"number_format": 3
	`
	whiteTextStr = `"font": {"color":"#ffffff"}`

	defaultStyle = `{
		` + defaultStyleStr + `, ` + roundOffToTwoStr + `
	}`

	greyColor = `
		"fill" :{"type":"pattern","color":["#bebbb8"],"pattern":1}
	`

	defaultGreyColor = `{
		` + defaultStyleStr + `, ` + roundOffToTwoStr +
		`, ` + greyColor + `, 
		"border": [
			{"type":"bottom","color":"FFFFFF","style":1},
			{"type":"left","color":"FFFFFF","style":1}
		]
		} `

	headerStyle = `{
		` + defaultStyleStr + `, ` + roundOffToTwoStr +
		`, "font": {"color":"#ffffff"},
		"fill" :{"type":"pattern","color":["#002060"],"pattern":1},
		"border": [
			{"type":"bottom","color":"FFFFFF","style":1},
			{"type":"left","color":"FFFFFF","style":1}
		]
	}`

	totalStyle = `{
		` + defaultStyleStr + `, ` + roundOffToTwoStr +
		`, "font": {"bold":true},
		"fill" :{"type":"pattern","color":["#D0CECE"],"pattern":1},
		"border": [
			{"type":"bottom","color":"FFFFFF","style":1},
			{"type":"left","color":"FFFFFF","style":1}
		]
	}`

	infoTextStyle = `{
		"alignment":{
			"horizontal":"center",
			"vertical":"center"
		},
		"font": {"bold":true, "underline": "single"}
	}`

	noColorHeaderStyle = `{
		` + defaultStyleStr + `, ` + roundOffToTwoStr +
		`, "font": {"color":"#333333"},
		"fill" :{"type":"pattern","pattern":1},
		"border": [
			{"type":"bottom","color":"#333333","style":1},
			{"type":"left","color":"#333333","style":1},
			{"type":"right","color":"#333333","style":1},
			{"type":"top","color":"#333333","style":1}
		]
	}`

	report4HeaderStyle = `{
		` + defaultStyleStr + `, ` + roundOffToTwoStr +
		`, "font": {"color":"#ffffff"},
		"fill" :{"type":"pattern","color":["#0B64A0"],"pattern":1},
		"border": [
			{"type":"bottom","color":"FFFFFF","style":1},
			{"type":"left","color":"FFFFFF","style":1}
		]
	}`

	roundOffToZero = `{ ` + defaultStyleStr + `, ` + commaNumberFormatStr + ` }`

	roundOffToZeroTotal = `{
		` + defaultStyleStr + `, ` + commaNumberFormatStr +
		`, "font": {"bold":true},
		"fill" :{"type":"pattern","color":["#CCE5FF"],"pattern":1},
		"border": [
			{"type":"bottom","color":"FFFFFF","style":1},
			{"type":"left","color":"FFFFFF","style":1}
		]
	}`
)

// GetStyle returns style ID for provided file
func GetStyle(f *excelize.File, s style) int {
	var sty int
	switch s {
	case HeaderSty:
		sty = getOrCreateStyle(f, s, headerStyle)
	case TotalSty:
		sty = getOrCreateStyle(f, s, totalStyle)
	case InfoTextSty:
		sty = getOrCreateStyle(f, s, infoTextStyle)
	case DefaultSty:
		sty = getOrCreateStyle(f, s, defaultStyle)
	case RoundOffToZeroSty:
		sty = getOrCreateStyle(f, s, roundOffToZero)
	case RoundOffToZeroTotalSty:
		sty = getOrCreateStyle(f, s, roundOffToZeroTotal)
	case Report4HeaderSty:
		sty = getOrCreateStyle(f, s, report4HeaderStyle)
	case NoColorHeaderSty:
		sty = getOrCreateStyle(f, s, noColorHeaderStyle)
	case DefaultGreyColor:
		sty = getOrCreateStyle(f, s, defaultGreyColor)

	}
	f.SetDefaultFont("Arial")
	return sty
}

func getOrCreateStyle(f *excelize.File, s style, styleStr string) int {
	var err error
	sty, ok := styles[f][s]
	if !ok {
		sty, err = f.NewStyle(styleStr)
		if err != nil {
			fmt.Println("Styling Error: ", err)
		}
	}
	if len(styles[f]) == 0 {
		styles[f] = make(map[style]int)
	}
	styles[f][s] = sty
	return sty
}

// SetActiveSheetAndDeleteDefaultSheet sets active sheet for the given sheet and remove the default sheet
func SetActiveSheetAndDeleteDefaultSheet(f *excelize.File, sheetName string) {
	f.DeleteSheet("Sheet1")
	// Set active sheet of the workbook.
	f.SetActiveSheet(f.GetSheetIndex(sheetName))
}
