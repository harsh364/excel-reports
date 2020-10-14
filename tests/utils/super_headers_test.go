package utils_test

import (
	"reflect"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"

	"example.com/emailreports/utils"
)

func TestGetPosition(t *testing.T) {
	t.Parallel()
	want := 1
	got := utils.GetPosition(want)
	if got != want {
		// t.Errorf("\n\n ===> TestGetPosition failed <===\n\n got  ===> %v \n\n want ===> %v \n\n", got, want)
		t.Error()
	}
}

func TestPositionIncr(t *testing.T) {
	t.Parallel()
	got := 1
	ip := 2
	want := 3
	op := utils.PositionIncr(&got, ip)
	if got != want {
		// t.Errorf("\n\n ===> TestPositionIncr failed <===\n\n got  ===> %v \n\n want ===> %v \n\n", got, want)
		t.Error()
	}
	if ip != op {
		// t.Errorf("\n\n ===> TestPositionIncr failed <===\n\n got  ===> %v \n\n want ===> %v \n\n", got, want)
		t.Error()
	}
}

var pos = 1
var addHeadersTestcases = []struct {
	name string
	ip   []utils.SuperHeader
	op   []string
}{
	{
		name: "Empty super headers",
		ip:   []utils.SuperHeader{},
		op:   []string{},
	},
	{
		name: "Valid case",
		ip: []utils.SuperHeader{
			{Name: "First", Position: utils.GetPosition(pos), RepeatFor: utils.PositionIncr(&pos, 1)},
			{Name: "Second", Position: utils.GetPosition(pos), RepeatFor: utils.PositionIncr(&pos, 4)},
		},
		op: []string{"First", "Second", "Second", "Second", "Second"},
	},
}

func TestAddSuperHeaderToSheet(t *testing.T) {
	t.Parallel()
	for _, tc := range addHeadersTestcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			row := 1
			f := excelize.NewFile()
			utils.AddSuperHeaderToSheet(tc.ip, f, "Sheet1", &row, utils.DefaultSty)
			rows, err := f.Rows("Sheet1")
			if err != nil {
				t.Error(err)
			}
			want := tc.op
			for rows.Next() {
				got, err := rows.Columns()
				if err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("\n\n ===> TestAddSuperHeaderToSheet failed for test case named *** %v *** <===\n\n got  ===> %#v \n\n want ===> %#v \n\n", tc.name, got, want)
					t.Error()
				}
			}
		})
	}

}
