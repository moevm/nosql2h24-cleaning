package types

import (
	"net/url"
	"time"
)

type TimeInterval struct {
	Begin *time.Time `json:"begin"`
	End   *time.Time `json:"end"`
}

func getTimePointer(query url.Values, key string) *time.Time {
	value, err := time.Parse(time.RFC3339, query.Get(key))
	if err != nil {
		return nil
	}
	return &value
}

func NewTimeInterval(key string, query url.Values) TimeInterval {
	return TimeInterval{
		Begin: getTimePointer(query, key+"_begin"),
		End:   getTimePointer(query, key+"_end"),
	}
}

func (t TimeInterval) GetBegin() *time.Time {
	return t.Begin
}

func (t TimeInterval) GetEnd() *time.Time {
	return t.End
}
