package utils

// Column is a type to store excel sheet columns
type Column int

// ColumnRowValue is a type which contains column and its corresponding row value
type ColumnRowValue struct {
	Column
	InsertionColumn Column
	RowValue        string
	Row             map[Column]interface{}
	ExtraColumns    Columns
}

// Columns is type for array of columns
type Columns []Column

const (
	// DefaultColumn is column which is set by golang as default when Column is initialized as a field
	DefaultColumn Column = iota
	// NoneColumn is column used for grouping when there is no grouping
	NoneColumn
	// FatherAgeColumn is column for Father's Age
	FatherAgeColumn
	// MotherAgeColumn is column for Mother's Age
	MotherAgeColumn
	// RaceColumn is column for race
	RaceColumn
	// AvgWeightolumn is column for weight
	AvgWeightolumn
)
