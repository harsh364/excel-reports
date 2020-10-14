package utils

import (
	"fmt"
	"strings"
	"time"
)

// StrListToStr joins list of string using comma required for BigQuery
// Input => ["abCd", "DEfg"]
// Output  => " \"abcd\", \"defg\" "
func StrListToStr(arr []string) string {
	return strings.ToLower("\"" + strings.Join(arr[:], "\",\"") + "\"")
}

// IntListToStr joins list of int using comma required for BigQuery
// Input => [1, 2, 3]
// Output  => "1, 2, 3"
func IntListToStr(intarr interface{}) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(intarr)), ","), "[]")
}

// MnthYr return map of year and months array from startDate and endDate
// Input => startDate = "2019-11-01", endDate = "2020-02-01"
// Output => { 2019: [11, 12], 2020: [1,2] }
func MnthYr(startDate, endDate time.Time) map[int][]int {
	var mnthYr = make(map[int][]int)
	endYr, endTimeMth, _ := endDate.Date()
	stYr, stTimeMth, _ := startDate.Date()
	stMth := int(stTimeMth)
	endMth := int(endTimeMth)
	for stYr < endYr || (stYr == endYr && stMth <= endMth) {
		mnthArr := mnthYr[stYr]
		if mnthArr == nil {
			mnthArr = []int{stMth}
		} else {
			mnthArr = append(mnthArr, stMth)
		}
		mnthYr[stYr] = mnthArr
		stMth++
		if stMth > 12 {
			stMth = 1
			stYr++
		}
	}
	return mnthYr
}
