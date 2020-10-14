package schema

import (
	"sort"

	sharedUtils "example.com/emailreports/reports/shared/utils"
)

// RowMap is a map with key as `string` as `Row Name` and value as `Row`
// Example => RowMap{"First Row": Row{ Column(0): "Channel Name", Column(1): 2.02 }}
type RowMap map[string]Row

// findOrCreate returns Row by finding it or returns new Row with the name provided
func (rowMap RowMap) findOrCreate(name string) Row {
	row, present := rowMap[name]
	if !present {
		row = Row{sharedUtils.RaceColumn: name}
	}
	return row
}

func (rowMap RowMap) findOrCreate2(name string, initialRow Row) Row {
	row, present := rowMap[name]
	if !present {
		row = initialRow
	}
	return row
}

// ConvertToTwoDecimal the RowMap numbers to two decimal
func (rowMap RowMap) ConvertToTwoDecimal() RowMap {
	newRowMap := RowMap{}
	for k, v := range rowMap {
		newRowMap[k] = v.ConvertToTwoDecimal()
	}
	return newRowMap
}

// Equal compares two RowMaps
func (rowMap RowMap) Equal(rowMap2 RowMap) bool {
	if &rowMap == &rowMap2 {
		return true
	}
	if len(rowMap) != len(rowMap2) {
		return false
	}
	for k, row := range rowMap {
		if !row.Equal(rowMap2[k]) {
			return false
		}
	}
	return true
}

// Sort return sorted array of rows
func (rowMap RowMap) Sort() []Row {
	var keys []string
	for k := range rowMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	sortedRowMap := make([]Row, len(rowMap))
	for i, k := range keys {
		sortedRowMap[i] = rowMap[k]
	}
	return sortedRowMap
}
