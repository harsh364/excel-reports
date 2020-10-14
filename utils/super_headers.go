package utils

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// SuperHeader is struct which has `name` i.e the name of the heder, `position` i.e of the column, and
// `RepeatFor` i.e how many consecutive number of columns to be filled with the header name
type SuperHeader struct {
	Name      string
	Position  int
	RepeatFor int
}

// ConvertToInterface return array name of SuperHeaders using reperatFor
func convertToInterface(sh []SuperHeader) []interface{} {
	arr := make([]interface{}, 0)
	for _, s := range sh {
		for i := 1; i <= s.RepeatFor; i++ {
			arr = append(arr, s.Name)
		}
	}
	return arr
}

// GetPosition returns updated position
func GetPosition(pos int) int {
	return pos
}

// PositionIncr increments the position pointer i.e. `posPtr` by provided value i.e. `val`.
func PositionIncr(posPtr *int, val int) int {
	*posPtr += val
	return val
}

// AddSuperHeaderToSheet add super-headers to sheet and merge the common columns
func AddSuperHeaderToSheet(superHeaders []SuperHeader, f *excelize.File, sheet string, row *int, sty style) {
	if len(superHeaders) == 0 {
		return
	}
	sheaders := convertToInterface(superHeaders)
	startPos, _ := excelize.ColumnNumberToName(superHeaders[0].Position)
	SetRow(f, sheet, startPos, row, sheaders, sty)
	mergeSuperHeaders(f, sheet, (*row)-1, superHeaders)
}

func mergeSuperHeaders(f *excelize.File, sheet string, row int, vals []SuperHeader) {
	for _, v := range vals {
		if v.RepeatFor > 1 {
			scol, _ := excelize.ColumnNumberToName(v.Position)
			srow := strconv.Itoa(row)
			ecol, _ := excelize.ColumnNumberToName(v.Position + v.RepeatFor - 1)
			erow := strconv.Itoa(row)
			f.MergeCell(sheet, scol+srow, ecol+erow)
		}
	}
}
