package utils_test

import (
	"reflect"
	"testing"

	"example.com/emailreports/utils"
)

var testCaseSortKeysArray = []struct {
	name string
	ip   map[string]interface{}
	op   []string
}{
	{
		name: "Default",
		ip: map[string]interface{}{
			"first":  true,
			"second": true,
			"third":  true,
		},
		op: []string{"first", "second", "third"},
	},
}

func TestSortKeysArray(t *testing.T) {
	t.Parallel()
	for _, tc := range testCaseSortKeysArray {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := utils.SortKeysArray(tc.ip)
			want := tc.op
			if !reflect.DeepEqual(got, want) {
				t.Errorf("\n\n ===> TestSortKeysArray failed for test case named *** %v *** <===\n\n got  ===> %v \n\n want ===> %v \n\n", tc.name, got, want)
			}
		})
	}
}
