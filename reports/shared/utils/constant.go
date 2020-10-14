package utils

// Average constant
const Average = "Average"

// RaceWiseWeight is a constant for Race Wise Weight tag
const RaceWiseWeight = "Race Wise Weight"

// Filter is type for filters
type Filter int

const (
	// Natality filter
	Natality Filter = iota + 1
)

// Group is a type for grouping columns
type Group string

const (
	// Group1 is a group name
	Group1 Group = "group1"
	// Group2 is a group name
	Group2 Group = "group2"
	// Group3 is a group name
	Group3 Group = "group3"
)
