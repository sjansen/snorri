package rrule

import (
	"time"

	"github.com/teambition/rrule-go"
)

type Frequency int

const (
	DAILY   = Frequency(rrule.DAILY)
	WEEKLY  = Frequency(rrule.WEEKLY)
	MONTHLY = Frequency(rrule.MONTHLY)
	YEARLY  = Frequency(rrule.YEARLY)
)

type WeekDay int

const (
	MO = WeekDay(1)
	TU = WeekDay(2)
	WE = WeekDay(3)
	TH = WeekDay(4)
	FR = WeekDay(5)
	SA = WeekDay(6)
	SU = WeekDay(7)
)

type WeekDayNum struct {
	WeekDay WeekDay
	Num     int
}

type Options struct {
	Frequency  Frequency
	Begin      time.Time
	Until      time.Time
	Count      int
	Interval   int
	WeekStart  WeekDay
	ByDay      []WeekDayNum
	ByMonthDay []int
	ByYearDay  []int
	ByWeekNo   []int
	ByMonth    []int
}

func toDate(dt time.Time) time.Time {
	return time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
}

func toDay(x WeekDay) rrule.Weekday {
	switch x {
	case MO:
		return rrule.MO
	case TU:
		return rrule.TU
	case WE:
		return rrule.WE
	case TH:
		return rrule.TH
	case FR:
		return rrule.FR
	case SA:
		return rrule.SA
	case SU:
		return rrule.SU
	}
	return rrule.Weekday{}
}

func toWeekday(x WeekDayNum) rrule.Weekday {
	switch x.WeekDay {
	case MO:
		return rrule.MO.Nth(x.Num)
	case TU:
		return rrule.TU.Nth(x.Num)
	case WE:
		return rrule.WE.Nth(x.Num)
	case TH:
		return rrule.TH.Nth(x.Num)
	case FR:
		return rrule.FR.Nth(x.Num)
	case SA:
		return rrule.SA.Nth(x.Num)
	case SU:
		return rrule.SU.Nth(x.Num)
	}
	return rrule.Weekday{}
}
