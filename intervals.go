package metha

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
)

// Interval represents a span of time.
type Interval struct {
	Begin time.Time
	End   time.Time
}

// String formats the interval.
func (iv Interval) String() string {
	return fmt.Sprintf("[%s--%s]", iv.Begin, iv.End)
}

// MonthlyIntervals segments a given interval into montly chunks.
func (iv Interval) MonthlyIntervals() []Interval {
	var ivals []Interval
	start := iv.Begin
	for {
		if start.After(iv.End) {
			break
		}
		end := now.New(start).EndOfMonth()
		if end.After(iv.End) {
			ivals = append(ivals, Interval{Begin: start, End: iv.End})
			break
		}
		ivals = append(ivals, Interval{Begin: start, End: end})
		start = now.New(start.AddDate(0, 1, 0)).BeginningOfMonth()
	}
	return ivals
}

// DailyIntervals segments a given interval into daily chunks.
func (iv Interval) DailyIntervals() []Interval {
	var ivals []Interval
	start := iv.Begin
	for {
		if start.After(iv.End) {
			break
		}
		end := now.New(start).EndOfDay()
		if end.After(iv.End) {
			ivals = append(ivals, Interval{Begin: start, End: end})
			break
		}
		ivals = append(ivals, Interval{Begin: start, End: end})
		start = now.New(start.AddDate(0, 0, 1)).BeginningOfDay()
	}
	return ivals
}
