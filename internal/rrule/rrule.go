package rrule

import (
	"time"

	"github.com/teambition/rrule-go"
)

type RRule struct {
	base *rrule.RRule
}

func New(o *Options) (*RRule, error) {
	ro := rrule.ROption{
		Freq:      rrule.Frequency(o.Frequency),
		Dtstart:   toDate(o.Begin),
		Interval:  o.Interval,
		Wkst:      toDay(o.WeekStart),
		Count:     o.Count,
		Until:     toDate(o.Until),
		Byweekday: make([]rrule.Weekday, len(o.ByDay)),
	}
	for idx, wd := range o.ByDay {
		ro.Byweekday[idx] = toWeekday(wd)
	}
	copy(ro.Bymonthday, o.ByMonthDay)
	copy(ro.Byyearday, o.ByYearDay)
	copy(ro.Byweekno, o.ByWeekNo)
	copy(ro.Bymonth, o.ByMonth)

	r, err := rrule.NewRRule(ro)
	if err != nil {
		return nil, err
	}

	rrule := &RRule{
		base: r,
	}

	return rrule, nil
}

func (r *RRule) All() []time.Time {
	return r.base.All()
}

func (r *RRule) String() string {
	return r.base.String()
}
