package schema

import "example.com/emailreports/reports/shared/utils"

// EntriesFetcher is a  entries fetcher interfcae
type EntriesFetcher interface {
	GetEntries(query string) Entries
}

// EntriesFetch is a fetcher to fetch  entries
type EntriesFetch struct{}

// GetEntries fetches  entries
func (f EntriesFetch) GetEntries(query string) Entries {
	return getEntries(query)
}

// ColumnFetcher is map for column and its data fetcher
type ColumnFetcher map[utils.Column]EntriesFetcher
