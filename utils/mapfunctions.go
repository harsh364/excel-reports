package utils

import "sort"

// SortKeysArray returns sorted keys array for map of key strings
func SortKeysArray(m map[string]interface{}) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
