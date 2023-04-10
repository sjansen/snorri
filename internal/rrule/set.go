package rrule

import (
	"encoding/json"
	"time"

	"github.com/teambition/rrule-go"
)

type Set struct {
	base *rrule.Set
}

func NewSet(r *RRule) *Set {
	s := &Set{
		base: &rrule.Set{},
	}
	s.base.RRule(r.base)
	return s
}

func (s *Set) ExDate(dt time.Time) *Set {
	s.base.ExDate(dt)
	return s
}

func (s *Set) RDate(dt time.Time) *Set {
	s.base.RDate(dt)
	return s
}

func (s *Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.base.String())
}

func (s *Set) UnmarshalJSON(b []byte) error {
	var tmp string
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	parsed, err := rrule.StrToRRuleSet(tmp)
	if err != nil {
		return err
	}

	s.base = parsed
	return nil
}
