package timesheet

import (
	"testing"
	"time"

	"github.com/gregoryv/asserter"
)

var (
	jan, feb *Sheet
)

func init() {
	var err error
	if jan, err = Load("assets/201801.timesheet"); err != nil {
		panic(err)
	}
	if feb, err = Load("assets/201802.timesheet"); err != nil {
		panic(err)
	}
}

func Test_making_a_report(t *testing.T) {
	report := NewReport()
	assert := asserter.New(t)
	assert(report != nil).Fail()
	assert(len(report.Sheets) == 0).Error("New reports should have no sheets")

	got, err := report.Append(jan)
	assert().Equals(got, 1)
	assert(err == nil).Error("Failed to add sheet")
}

func Test_summarize_reports(t *testing.T) {
	report := newTestReport()
	total := report.Reported()
	exp := 179*time.Hour + 30*time.Minute // jan
	exp += 155 * time.Hour                // feb
	assert := asserter.New(t)
	assert().Equals(total, exp)
}

func Test_tag_summary(t *testing.T) {
	r := newTestReport()
	tags := r.Tags()
	assert := asserter.New(t)
	assert().Equals(len(tags), 2)
}

func Test_find_in_report(t *testing.T) {
	r := newTestReport()
	period := "2018 January"
	sheet, err := r.FindByPeriod(period)
	assert := asserter.New(t)
	assert(err == nil).Fatal(err)
	assert().Equals(sheet.Period, period)

	_, err = r.FindByPeriod("whoops")
	assert(err != nil).Fail()
}

func newTestReport() (r *Report) {
	r = NewReport()
	r.Append(jan)
	r.Append(feb)
	return
}
