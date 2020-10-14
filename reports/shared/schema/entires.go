package schema

import (
	"fmt"
	"sort"

	"cloud.google.com/go/bigquery"
	sharedUtils "example.com/emailreports/reports/shared/utils"
	"example.com/emailreports/utils"
	"google.golang.org/api/iterator"
)

// Entries is a list of Entry
type Entries []Entry

// EntriesArr is a list of Entries
type EntriesArr []Entries

// EntriesGroups is group of  entries grouped by a column
type EntriesGroups map[interface{}]Entries

// =============================================================
// ============ Entries Methods/Functions ===============
// =============================================================

// Compare compares between two  entries
func (res Entries) Compare(i, j int) bool {
	return res[i].compare(res[j])
}

// GroupBy groups the  entries by the Column provided
func (res Entries) GroupBy(col sharedUtils.Column) EntriesGroups {
	groups := make(EntriesGroups)
	for _, re := range res {
		key := re.GetColumnValue(col)
		subRes := groups[key]
		groups[key] = append(subRes, re)
	}
	return groups
}

// Merge merges all the  entries to one  entry
func (res Entries) Merge(divideBy DivideBy) Entry {
	mergeRe := Entry{}
	for _, re := range res {
		mergeRe = mergeRe.Add(re)
		// fmt.Println(mergeRe)
	}
	return mergeRe
}

// SetColumn sets `SetColumn` for  entries
func (res Entries) SetColumn(col sharedUtils.Column) {
	for i := range res {
		res[i].Column = col
	}
}

// Equal checks equality for Entries
func (res Entries) Equal(newRes Entries) bool {
	if len(res) == len(newRes) {
		sort.Slice(res, res.Compare)
		sort.Slice(newRes, newRes.Compare)
		for i, re := range res {
			if !re.Equal(newRes[i]) {
				return false
			}
		}
		return true
	}
	return false
}

// GetEntries fetches  for channel by region or genre
func getEntries(query string) Entries {
	var bks Entries
	it := utils.BqQuery(query)
	for {
		// var bk Entry
		// err := it.Next(&bk)
		var m map[string]bigquery.Value
		err := it.Next(&m)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("Error in parsing Get Bookings: Schema: ", err)
		}
		bks = append(bks, GetEntry(m))
	}
	return bks
}

// =============================================================
// ============ EntriesArr Methods/Functions ============
// =============================================================

// Compare compares between two list of  entries
func (resArr EntriesArr) Compare(i, j int) bool {
	sort.Slice(resArr[i], resArr[i].Compare)
	sort.Slice(resArr[j], resArr[j].Compare)
	if len(resArr[i]) > 0 && len(resArr[j]) > 0 {
		return resArr[i][0].compare(resArr[j][0])
	}
	return false
}

// Equal checks equality for EntriesArr's
func (resArr EntriesArr) Equal(newResArr EntriesArr) bool {
	if len(resArr) == len(newResArr) {
		sort.Slice(resArr, resArr.Compare)
		sort.Slice(newResArr, newResArr.Compare)
		for i, res := range resArr {
			if !res.Equal(newResArr[i]) {
				return false
			}
		}
		return true
	}
	return false
}

// ================================================================
// ============ EntriesGroups Methods/Functions ============
// ================================================================

// Sort sorts all the slices of  entries in EntriesGroups
func (reGrp EntriesGroups) Sort() EntriesGroups {
	for _, res := range reGrp {
		sort.Slice(res, res.Compare)
	}
	return reGrp
}
