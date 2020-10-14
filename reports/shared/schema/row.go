package schema

import (
	"fmt"
	"math"

	sharedUtils "example.com/emailreports/reports/shared/utils"
	"example.com/emailreports/utils"
)

// Row is type for row of sheet
// It is a map with key as `Column` and Value as `empty interface`
// Example => Row{ Column(0): "Channel Name", Column(1): 2.02 }
type Row map[sharedUtils.Column]interface{}

//ToArray converts rowColumnsMapping to Array of empty interface
func (r Row) ToArray(cols sharedUtils.Columns) []interface{} {
	var values []interface{}
	for _, c := range cols {
		values = append(values, r[c])
	}
	return values
}

func (r Row) rowCopy() Row {
	newRow := Row{}
	for c, val := range r {
		newRow[c] = val
	}
	return newRow
}

// ConvertToTwoDecimal the Row numbers to two decimal
func (r Row) ConvertToTwoDecimal() Row {
	newRow := Row{}
	for k, v := range r {
		switch v.(type) {
		case float32:
			newRow[k] = utils.RoundFloat(float64(v.(float32)))
		case float64:
			newRow[k] = utils.RoundFloat(v.(float64))
		case int:
			newRow[k] = utils.RoundFloat(float64(v.(int)))
		default:
			newRow[k] = v
		}
	}
	return newRow
}

// Equal compares two Rows
func (r Row) Equal(r2 Row) bool {
	// if &r == &r2 {
	// 	return true
	// }
	if len(r) != len(r2) {
		return false
	}
	for k, v := range r {
		allowedError := 0.1
		if val, ok := v.(float64); ok && val > 1000 {
			allowedError = 10.0
		}
		switch r2[k].(type) {
		case float32:
			errorVal := v.(float64) - float64(r2[k].(float32))
			if math.Abs(errorVal) > allowedError {
				fmt.Println("float32", v.(float64), r2[k].(float32), errorVal)
				return false
			}
		case float64:
			errorVal := v.(float64) - r2[k].(float64)
			if math.Abs(errorVal) > allowedError {
				fmt.Println("float64", v.(float64), r2[k].(float64), errorVal)
				return false
			}
		case int:
			errorVal := v.(float64) - float64(r2[k].(int))
			if math.Abs(errorVal) > allowedError {
				fmt.Println("int", v.(float64), r2[k].(int), errorVal)
				return false
			}
		default:
			if r2[k] != v {
				fmt.Println("default", v, r2[k])
				return false
			}
		}
	}
	return true
}

//setPercent to calculate the difference between the get columns and set value in set column
func (r Row) setPercent(numerator, denominator, setColumn sharedUtils.Column) {
	var num float64
	var deno = float64(1)
	if value, present := r[numerator]; present {
		num = value.(float64)
	}
	if value, present := r[denominator]; present && value.(float64) > 0 {
		deno = value.(float64)
	}
	r[setColumn] = (num / deno) * 100
}
