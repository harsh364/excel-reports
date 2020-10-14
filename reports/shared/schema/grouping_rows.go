package schema

import (
	"fmt"
	"sort"
	"strconv"

	"cloud.google.com/go/civil"
	sharedUtils "example.com/emailreports/reports/shared/utils"
)

// GroupingRow is structure to hold the related group of rows
type GroupingRow struct {
	Name            string
	Parent          Row
	Children        RowMap
	SubGroupingRows *GroupingRows
}

// IsEmpty return true if the grouping row is empty
func (grpRow *GroupingRow) IsEmpty() bool {
	return grpRow.Equal(GroupingRow{})
}

// GroupingRows is a map holding groupingRows
// consisting of key as RaceColumn and value as GroupingRow
type GroupingRows map[string]GroupingRow

// Grouping is used for grouping
func (grpRows *GroupingRows) Grouping(res Entries, attr EntryAttribute, groupers []sharedUtils.ColumnRowValue, commonRow Row, transposeMethod string, divideBy DivideBy) {
	if len(groupers) == 2 {
		switch transposeMethod {
		case "transpose":
			grpRows.Transpose(res, groupers[0], groupers[1], attr, divideBy)
		case "transpose2":
			grpRows.transpose2(res, groupers[0], groupers[1], attr, commonRow.rowCopy(), divideBy)
		default:
			panic("No such method error:" + transposeMethod)
		}
		return
	}
	revGrps := EntriesGroups{}
	col := groupers[0].Column
	if col == sharedUtils.NoneColumn {
		// No Grouping
		revGrps = EntriesGroups{groupers[0].RowValue: res}
		col = sharedUtils.RaceColumn
	} else {
		// Grouping by parents
		revGrps = res.GroupBy(col)
	}
	insertIntoColumn := groupers[0].InsertionColumn
	if insertIntoColumn == sharedUtils.DefaultColumn {
		insertIntoColumn = col
	}

	for key, grpRes := range revGrps {
		newCommonRow := commonRow.rowCopy()
		newCommonRow[col] = key
		newCommonRow[insertIntoColumn] = key
		key = getKeyString(key)
		subGrp := grpRows.findOrCreate2(key.(string), newCommonRow)
		// for k, v := range commonRow {
		// 	subGrp.Parent[k] = v
		// }
		re := grpRes.Merge(divideBy)
		for _, extraCol := range groupers[0].ExtraColumns {
			newCommonRow[extraCol] = re.GetColumnValue(extraCol)
		}
		subGrp.Parent[re.Column] = re.GetAttributeValue(attr)
		if subGrp.SubGroupingRows == nil {
			newGrps := GroupingRows{}
			subGrp.SubGroupingRows = &newGrps
		}
		subGrp.SubGroupingRows.Grouping(grpRes, attr, groupers[1:], newCommonRow.rowCopy(), transposeMethod, divideBy)
		(*grpRows)[key.(string)] = subGrp
	}
}

func (grpRows GroupingRows) transpose2(res Entries, parentGroupColumnRowValue, childGroupColumnRowValue sharedUtils.ColumnRowValue, attr EntryAttribute, commonRow Row, divideBy DivideBy) {
	var chGrp EntriesGroups
	parentCol := parentGroupColumnRowValue.Column
	parentRowVal := parentGroupColumnRowValue.RowValue
	if parentCol == sharedUtils.NoneColumn {
		// No Grouping
		chGrp = EntriesGroups{parentRowVal: res}
		parentCol = sharedUtils.RaceColumn
		// if parentCol == sharedUtils.NoneColumn {
		// }
	} else {
		// Grouping by parents
		chGrp = res.GroupBy(parentCol)
	}
	insertIntoColumn := parentGroupColumnRowValue.InsertionColumn
	if insertIntoColumn == sharedUtils.DefaultColumn {
		insertIntoColumn = parentCol
	}
	for ch, grpRes := range chGrp {
		commonRow := commonRow.rowCopy()
		chStr := getKeyString(ch)
		commonRow[insertIntoColumn] = chStr
		// commonRow[parentCol] = chStr
		re := grpRes.Merge(divideBy)
		for _, extraCol := range parentGroupColumnRowValue.ExtraColumns {
			commonRow[extraCol] = re.GetColumnValue(extraCol)
		}
		grpRow := grpRows.findOrCreate2(chStr, commonRow)
		for k, v := range parentGroupColumnRowValue.Row {
			grpRow.Parent[k] = v
		}

		grpRow.Parent[re.Column] = re.GetAttributeValue(attr)

		// Group Children
		childCol := childGroupColumnRowValue.Column
		childInsertIntoColumn := childGroupColumnRowValue.InsertionColumn
		if childInsertIntoColumn == sharedUtils.DefaultColumn {
			childInsertIntoColumn = childCol
		}
		if childCol != sharedUtils.NoneColumn {
			subGrps := grpRes.GroupBy(childCol)
			for subGrpKey, res := range subGrps {
				commonRow := commonRow.rowCopy()
				subGrpKeyStr := getKeyString(subGrpKey)
				commonRow[childInsertIntoColumn] = subGrpKeyStr
				// commonRow[childCol] = subGrpKeyStr
				re := res.Merge(divideBy)
				for _, extraCol := range childGroupColumnRowValue.ExtraColumns {
					commonRow[extraCol] = re.GetColumnValue(extraCol)
				}
				row := grpRow.Children.findOrCreate2(subGrpKeyStr, commonRow)
				for k, v := range commonRow {
					row[k] = v
				}

				row[re.Column] = re.GetAttributeValue(attr)
				grpRow.Children[subGrpKeyStr] = row
			}
		}
		grpRows[chStr] = grpRow
	}
}

// Transpose transposes `Entries` array columns to `GroupingRow` array rows
func (grpRows GroupingRows) Transpose(res Entries, parentGroupColumnRowValue, childGroupColumnRowValue sharedUtils.ColumnRowValue, attr EntryAttribute, divideBy DivideBy) {
	var chGrp EntriesGroups
	if parentGroupColumnRowValue.Column == sharedUtils.NoneColumn {
		// No Grouping
		chGrp = EntriesGroups{parentGroupColumnRowValue.RowValue: res}
	} else {
		// Grouping by parents
		chGrp = res.GroupBy(parentGroupColumnRowValue.Column)
	}
	for ch, grpRes := range chGrp {
		chStr := ch.(string)
		re := grpRes.Merge(divideBy)
		grpRow := grpRows.findOrCreate(chStr)
		grpRow.Parent[re.Column] = re.GetAttributeValue(attr)
		// Group Children
		if childGroupColumnRowValue.Column != sharedUtils.NoneColumn {
			subGrps := grpRes.GroupBy(childGroupColumnRowValue.Column)
			for subGrpKey, res := range subGrps {
				subGrpKeyStr := getKeyString(subGrpKey)

				re := res.Merge(divideBy)
				row := grpRow.Children.findOrCreate(subGrpKeyStr)
				row[re.Column] = re.GetAttributeValue(attr)
				grpRow.Children[subGrpKeyStr] = row
			}
		}
		grpRows[chStr] = grpRow
	}
}

// findOrCreate returns GroupingRow by finding it or returns new GroupingRow with the name provided
func (grpRows *GroupingRows) findOrCreate(name string) GroupingRow {
	grpRow, present := (*grpRows)[name]
	if !present {
		grpRow = GroupingRow{
			Name:     name,
			Parent:   Row{sharedUtils.RaceColumn: name},
			Children: make(RowMap),
		}
	}
	return grpRow
}

func (grpRows *GroupingRows) findOrCreate2(name string, initalRow Row) GroupingRow {
	grpRow, present := (*grpRows)[name]
	if !present {
		grpRow = GroupingRow{
			Name:     name,
			Parent:   initalRow,
			Children: make(RowMap),
		}
	}
	return grpRow
}

// ExtractTotalRow removes TotalGroupingRow from GroupingRows and returns both the TotalGroupingRow and remaining GroupingRows
func (grpRows GroupingRows) ExtractTotalRow(totalRowName string) (GroupingRow, GroupingRows) {
	var totalGrpRow GroupingRow
	extractedGrpRows := make(GroupingRows)
	for k, v := range grpRows {
		if k == totalRowName {
			totalGrpRow = v
			continue
		}
		extractedGrpRows[k] = v
	}
	return totalGrpRow, extractedGrpRows
}

// Equal compares two GroupingRows
func (grpRows GroupingRows) Equal(grpRows2 GroupingRows) bool {
	// if &grpRows == &grpRows2 {
	// 	return true
	// }
	if len(grpRows) != len(grpRows2) {
		return false
	}
	for k, grpRow := range grpRows {
		if !grpRow.Equal(grpRows2[k]) {
			return false
		}
	}
	return true
}

// Equal compares two GroupingRow
func (grpRow GroupingRow) Equal(grpRow2 GroupingRow) bool {
	// if &grpRow == &grpRow2 {
	// 	return true
	// }
	if grpRow.Name != grpRow2.Name ||
		!grpRow.Parent.Equal(grpRow2.Parent) ||
		!grpRow.Children.Equal(grpRow2.Children) ||
		(grpRow.SubGroupingRows != nil && grpRow2.SubGroupingRows == nil) ||
		(grpRow.SubGroupingRows == nil && grpRow2.SubGroupingRows != nil) ||
		(grpRow.SubGroupingRows != nil && grpRow2.SubGroupingRows != nil && !grpRow.SubGroupingRows.Equal(*grpRow2.SubGroupingRows)) {
		return false
	}
	return true
}

// ConvertToTwoDecimal the GroupingRows numbers to two decimal
func (grpRows GroupingRows) ConvertToTwoDecimal() GroupingRows {
	newGrpRows := GroupingRows{}
	for k, grpRow := range grpRows {
		newGrpRows[k] = grpRow.ConvertToTwoDecimal()
	}
	return newGrpRows
}

// ConvertToTwoDecimal the GroupingRow numbers to two decimal
func (grpRow GroupingRow) ConvertToTwoDecimal() GroupingRow {
	newGroupingRow := GroupingRow{}
	newGroupingRow.Parent = grpRow.Parent.ConvertToTwoDecimal()
	newGroupingRow.Children = grpRow.Children.ConvertToTwoDecimal()
	if grpRow.SubGroupingRows != nil {
		newSubGroupingRows := grpRow.SubGroupingRows.ConvertToTwoDecimal()
		newGroupingRow.SubGroupingRows = &newSubGroupingRows
	}
	return newGroupingRow
}

// Sort return sorted array of grouping rows
func (grpRows GroupingRows) Sort() []GroupingRow {
	var keys []string
	for k := range grpRows {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	sortedGrpRows := make([]GroupingRow, len(grpRows))
	for i, k := range keys {
		sortedGrpRows[i] = grpRows[k]
	}
	return sortedGrpRows
}

// CalculatePercent calculates the percent of the get columns passed and set the updated value in setColumn
func (grpRows *GroupingRows) CalculatePercent(getColumn1, getColumn2, setColumn sharedUtils.Column) {
	for _, grpRow := range *grpRows {
		grpRow.Parent.setPercent(getColumn1, getColumn2, setColumn)
		for _, childRow := range grpRow.Children {
			childRow.setPercent(getColumn1, getColumn2, setColumn)
		}
		if grpRow.SubGroupingRows != nil {
			grpRow.SubGroupingRows.CalculatePercent(getColumn1, getColumn2, setColumn)
		}
	}
}

func getKeyString(key interface{}) string {
	var keyStr string
	switch key.(type) {
	case civil.Date:
		keyStr = key.(civil.Date).String()
	case string:
		keyStr = key.(string)
	case int:
		keyStr = strconv.Itoa(key.(int))
	default:
		err := fmt.Errorf("No type define for: %T", key)
		panic(err)
	}
	return keyStr
}
