package schema

import (
	sharedSchema "example.com/emailreports/reports/shared/schema"
)

// WeightByParentsAgeKey is a struct for grouping  entries
type WeightByParentsAgeKey struct{}

// GetKey is used for grouping WeightByParentsAgeKey  entries
func (k WeightByParentsAgeKey) GetKey(re sharedSchema.Entry) sharedSchema.GroupingKey {
	return sharedSchema.GroupingKey{
		FatherAge: re.FatherAge,
		MotherAge: re.MotherAge,
	}
}

// RaceWiseWeightKey is a struct for grouping  entries
type RaceWiseWeightKey struct{}

// GetKey is used for grouping SummaryKey  entries
func (k RaceWiseWeightKey) GetKey(re sharedSchema.Entry) sharedSchema.GroupingKey {
	return sharedSchema.GroupingKey{
		Race:      re.Race,
		FatherAge: re.FatherAge,
	}
}
