package types

import (
	"net/url"
	"strconv"
)

// NumberRange represents a range of numbers.
type NumberRange struct {
	Min *int `json:"min"`
	Max *int `json:"max"`
}

func getIntPointer(query url.Values, key string) *int {
	value, err := strconv.Atoi(query.Get(key))
	if err != nil {
		return nil
	}
	return &value
}

func NewNumberRange(key string, query url.Values) NumberRange {
	return NumberRange{
		Min: getIntPointer(query, key+"_min"),
		Max: getIntPointer(query, key+"_max"),
	}
}

// GetMin returns the minimum pointer to value of the range.
func (nr NumberRange) GetMin() *int {
	return nr.Min
}

// GetMax returns the maximum pointer to value of the range.
func (nr NumberRange) GetMax() *int {
	return nr.Max
}
