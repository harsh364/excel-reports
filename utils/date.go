package utils

import (
	"strconv"
	"time"

	"cloud.google.com/go/civil"
)

// HeaderDateLayout is a layout for date header
const HeaderDateLayout = "02/01/2006"

// AsOfDateLayout is a layout for as of date to display in excel
const AsOfDateLayout = "January 2, 2006"

// AsOnDateReport2Layout is a layout for as of date to display in excel
// const AsOnDateReport2Layout = "01/02/2006: 15:04:05"
const AsOnDateReport2Layout = AsOfDateLayout

// CombinedDateFormat is a Combined Report layout for as of date to display in excel
const CombinedDateFormat = "2006-01-02"

const dateFormat = "2006-1-2"

// Year returns year
type Year int

// Month returns month
type Month time.Month

// YearMonthMap is a structure of year pointing to array of months
// Ex => { 2019: [11, 12], 2020: [1,2] }
type YearMonthMap map[Year][]Month

// DateLessThanEqula return true if the start date is less than or equal to end date else returns false
func DateLessThanEqula(startDate, endDate time.Time) bool {
	stYr, stMnt, stDay := startDate.Date()
	enYr, enMnt, enDay := endDate.Date()
	if stYr < enYr {
		return true
	}
	if stYr == enYr {
		if stMnt < enMnt {
			return true
		}
		if stMnt == enMnt {
			return stDay <= enDay
		}
	}
	return false
}

// SubtractYear subtracts 1 year from the time.Time (i.e. Date) with leap year adjustment
func SubtractYear(date time.Time) time.Time {
	_, mnt, day := date.Date()
	leapYearAdjustment := 0
	if mnt == time.February && day == 29 {
		leapYearAdjustment = 1
	}
	return date.AddDate(-1, 0, leapYearAdjustment)
}

// ToDateFromCivilDate converts to time.Time (i.e. Date)
func ToDateFromCivilDate(date civil.Date) time.Time {
	formatedDate, _ := time.Parse(dateFormat, strconv.Itoa(date.Year)+"-"+strconv.Itoa(int(date.Month))+"-"+strconv.Itoa(date.Day))
	return formatedDate
}

// ToStr converts time.Time to Date string in format "yyyy-mm-dd"
func ToStr(dateTime time.Time) string {
	year, month, day := dateTime.Date()
	return strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
}

// ToDate converts date string of type "yyyy-mm-dd" to date
func ToDate(dString string) time.Time {
	date, err := time.Parse(dateFormat, dString)
	if err != nil {
		panic(err)
	}
	return date
}
