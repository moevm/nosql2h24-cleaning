package types

// NumberRange represents a range of numbers.
type NumberRange struct {
	Min *int `json:"min"`
	Max *int `json:"max"`
}

// GetMin returns the minimum pointer to value of the range.
func (nr NumberRange) GetMin() *int {
	return nr.Min
}

// GetMax returns the maximum pointer to value of the range.
func (nr NumberRange) GetMax() *int {
	return nr.Max
}
