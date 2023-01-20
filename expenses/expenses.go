package expenses

import (
	"errors"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var r []Record

	for _, v := range in {
		if predicate(v) {
			r = append(r, v)
		}
	}

	return r
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	f := func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}

	return f
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	f := func(r Record) bool {
		return r.Category == c
	}

	return f
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var total float64

	r := Filter(in, ByDaysPeriod(p))

	for _, v := range r {
		total += v.Amount
	}

	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	errUnknown := errors.New("unknown category")

	r := Filter(in, ByCategory(c))
	if r == nil {
		return 0, errUnknown
	}

	total := TotalByPeriod(r, p)

	return total, nil
}
