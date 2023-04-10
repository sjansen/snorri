package main

import (
	"fmt"
	"time"

	"github.com/sjansen/snorri/internal/rrule"
)

func main() {
	r, _ := rrule.New(&rrule.Options{
		Frequency:  rrule.MONTHLY,
		Begin:      time.Date(2023, 4, 9, 0, 0, 0, 0, time.UTC),
		ByMonthDay: []int{6, 7, 8, 9, 10, 11, 12},
		ByDay: []rrule.WeekDayNum{
			{WeekDay: rrule.SU, Num: 2},
		},
		Count: 6,
	})

	fmt.Println(r.String())
	for _, t := range r.All() {
		fmt.Println(t)
	}
}
