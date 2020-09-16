package autodmn

import "time"

type Compare struct {
}

func NewCompare() Compare {
	return Compare{}
}

func (c Compare) CompareInt(value, target int, op ComparisonOperator) bool {
	switch op {
	case EQ:
		return value == target
	case LT:
		return value < target
	case GT:
		return value > target
	case LET:
		return value <= target
	case GET:
		return value >= target
	}
	return false
}

func (c Compare) CompareFloat64(value, target float64, op ComparisonOperator) bool {
	switch op {
	case EQ:
		return value == target
	case LT:
		return value < target
	case GT:
		return value > target
	case LET:
		return value <= target
	case GET:
		return value >= target
	}
	return false
}

func (c Compare) BetweenFloat64(value, min, max float64, op ComparisonOperator) bool {
	if value < min || value > max {
		return false
	}
	return true
}
func (c Compare) BetweenInt(value, min, max int, op ComparisonOperator) bool {
	if value < min || value > max {
		return false
	}
	return true
}

func (c Compare) Between(a, b int) bool {
	return a < b
}

func (c Compare) CompareIsAfterTime(value, target time.Time) bool {
	return value.After(target)
}

func (c Compare) CompareIsBeforeTime(value, target time.Time) bool {
	return value.Before(target)
}
