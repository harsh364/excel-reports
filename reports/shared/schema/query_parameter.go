package schema

import (
	"time"

	sharedUtils "example.com/emailreports/reports/shared/utils"
)

// QueryType define query parameter type
type QueryType int

const (
	// WeightQueryType is a Weight type query parameter
	WeightQueryType QueryType = iota
)

// Selection is a type for big query column selection
type Selection int

// Selections is an array of Selection
type Selections []Selection

const (
	// RaceSelection is selection for Race_name
	RaceSelection Selection = iota
	// FatherAgeSelection is selection for FatherAge
	FatherAgeSelection
	// MotherAgeSelection is selection for MotherAge
	MotherAgeSelection
)

// FilterClause is a struct which generates where clause in query
type FilterClause struct {
	Attribute    string   `json:"attribute"`
	Inclusion    bool     `json:"inclusion"`
	NullRequired bool     `json:"null_required"`
	Options      []string `json:"options"`
}

// FilterClauses is an array FilterClause
type FilterClauses []FilterClause

// Merge merges all the give FilterClauses into new FilterClauses
func Merge(fcsArr ...FilterClauses) FilterClauses {
	nfcs := FilterClauses{}
	for _, fcs := range fcsArr {
		nfcs = append(nfcs, fcs...)
	}
	return nfcs
}

// FilterType is the type of filter
type FilterType string

const (
	// CommonFilterType filter type
	CommonFilterType FilterType = "common"
	// WeightFilterType filter type
	WeightFilterType = "weight"
)

// FilterClauseMap is an map data-type i.e('weight') and  FilterClauses
type FilterClauseMap map[FilterType]FilterClauses

// QueryParameter is a structure to define input parameters for the query
type QueryParameter struct {
	Name       string
	Column     sharedUtils.Column
	ReportDate time.Time
	filter     sharedUtils.Filter
	Type       QueryType
	Selections
	FilterClauses
	Fetcher    EntriesFetcher
	FromClause func(QueryParameter, FilterClauses) string
}

// SetFilter return query parameter with filter value for filter provided
func (qp QueryParameter) SetFilter(f sharedUtils.Filter) QueryParameter {
	qp.filter = f
	return qp
}

// GetFilter returns filter for Query Parameter
func (qp QueryParameter) GetFilter() sharedUtils.Filter {
	if qp.filter == 0 {
		return defaultFilter()
	}
	return qp.filter
}

// FilterIsEmpty return true if no filter set
func (qp *QueryParameter) FilterIsEmpty() bool {
	return qp.filter == 0
}

// SetReportDate return query parameter with filter value for filter provided
func (qp QueryParameter) SetReportDate(d time.Time) QueryParameter {
	qp.ReportDate = d
	return qp
}

// SetColumn return query parameter with column value for column provided
func (qp QueryParameter) SetColumn(c sharedUtils.Column) QueryParameter {
	qp.Column = c
	return qp
}

// SetName return query parameter with name for name provided
func (qp QueryParameter) SetName(name string) QueryParameter {
	queryFor := ""
	switch qp.Type {
	case WeightQueryType:
		queryFor = " Weight"
	}
	qp.Name = name + queryFor + " Query"
	return qp
}

func defaultFilter() sharedUtils.Filter {
	return sharedUtils.Natality
}

// SetType return query parameter with provided Type value
func (qp QueryParameter) SetType(t QueryType) QueryParameter {
	qp.Type = t
	return qp
}

// AddSelection add selection to query parameter
func (qp QueryParameter) AddSelection(sel Selection) QueryParameter {
	qp.Selections = append(qp.Selections, sel)
	return qp
}

// AddFilterClause add filter to query parameter
func (qp QueryParameter) AddFilterClause(fil FilterClause) QueryParameter {
	qp.FilterClauses = append(qp.FilterClauses, fil)
	return qp
}

func (qp QueryParameter) hasSelection(sel Selection) bool {
	for _, s := range qp.Selections {
		if s == sel {
			return true
		}
	}
	return false
}
