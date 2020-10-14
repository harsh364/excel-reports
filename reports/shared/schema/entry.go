package schema

import (
	"fmt"

	"cloud.google.com/go/bigquery"
	sharedUtils "example.com/emailreports/reports/shared/utils"
)

// Entry is  on channel, genre or region cut
type Entry struct {
	Weight    float64 `bigquery:"weight_pounds,nullable" json:"weight_pounds"`
	Race      int     `bigquery:"child_race,nullable" json:"child_race"`
	FatherAge int     `bigquery:"father_age,nullable" json:"father_age"`
	MotherAge int     `bigquery:"mother_age,nullable" json:"mother_age"`
	sharedUtils.Column
}

// EntryAttribute type for Entry attribute
type EntryAttribute int

const (
	// Weight attribute of Entry
	Weight EntryAttribute = iota + 1
)

// ==============================================================
// =============== Entry Methods/Functions ===============
// ==============================================================

// GetColumnValue returns Column value
func (re Entry) GetColumnValue(col sharedUtils.Column) interface{} {
	var key interface{}
	switch col {
	case sharedUtils.RaceColumn:
		key = re.Race
	case sharedUtils.FatherAgeColumn:
		key = re.FatherAge
	case sharedUtils.MotherAgeColumn:
		key = re.MotherAge
	case sharedUtils.AvgWeightolumn:
		key = re.Weight
	default:
		err := fmt.Errorf("No  Entry attribute defined for column %T", col)
		panic(err)
	}
	return key
}

// GetAttributeValue returns value of the attribute attached to it
func (re Entry) GetAttributeValue(attr EntryAttribute) float64 {
	var val float64
	switch attr {

	case Weight:
		val = re.Weight
	default:
		panic("No  Entry attribute defined provided attribute")
	}
	return val
}

// Add adds the given  entry
func (re Entry) Add(newRe Entry) Entry {
	re.Race = newRe.Race
	re.FatherAge = newRe.FatherAge
	re.MotherAge = newRe.MotherAge
	re.Weight += newRe.Weight
	return re
}

// Equal return true if the  entry is equal with minor error
func (re Entry) Equal(newRe Entry) bool {
	if re.Race == newRe.Race &&
		re.FatherAge == newRe.FatherAge &&
		re.MotherAge == newRe.MotherAge &&
		re.Weight == newRe.Weight {
		return true
	}
	return false
}

func (re Entry) compare(resj Entry) bool {
	if re.Race < resj.Race {
		return true
	} else if re.Race == resj.Race {
		if re.FatherAge < resj.FatherAge {
			return true
		} else if re.FatherAge == resj.FatherAge {
			if re.MotherAge < resj.MotherAge {
				return true
			} else if re.MotherAge == resj.MotherAge {
				if re.Column < resj.Column {
					return true
				} else if re.Column == resj.Column {
					return re.Weight < resj.Weight

				}
			}
		}
	}

	return false
}

// GetEntry generates Entry from bigquery data
func GetEntry(data map[string]bigquery.Value) Entry {
	re := Entry{}
	for k, v := range data {
		switch k {
		case "child_race":
			val, ok := v.(int64)
			if ok {
				re.Race = int(val)
			}
		case "weight_pounds":
			val, ok := v.(float64)
			if ok {
				re.Weight = val
			}
		case "father_age":
			val, ok := v.(int64)
			if ok {
				re.FatherAge = int(val)
			}
		case "mother_age":
			val, ok := v.(int64)
			if ok {
				re.MotherAge = int(val)
			}
		}
	}
	return re
}
