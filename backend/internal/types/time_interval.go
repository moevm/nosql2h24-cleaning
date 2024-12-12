package types

import "time"

type TimeInterval struct {
	Begin *time.Time
	End   *time.Time
}

func (t TimeInterval) GetBegin() *time.Time {
	return t.Begin
}

func (t TimeInterval) GetEnd() *time.Time {
	return t.End
}
