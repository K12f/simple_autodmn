package autoaddmn

import (
	"time"
)

// 格式化日期，默认格式
const TimeFormat = "2006-01-02 15:04:05"

// 符号标识
type ComparisonOperator string

const (
	EQ      ComparisonOperator = "="
	GT      ComparisonOperator = ">"
	GET     ComparisonOperator = ">="
	LT      ComparisonOperator = "<"
	LET     ComparisonOperator = "<="
	BETWEEN ComparisonOperator = "in"
)

// 比较结构体
type Compare struct {
}

func NewCompare() Compare {
	return Compare{}
}

// 根据 操作符和值 比较两个int类型
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

// 根据 操作符和值 比较两个float类型
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

// 根据 输入的 float类型值，判断 数据是否在 [min max] 之间
func (c Compare) BetweenFloat64(value, min, max float64, op ComparisonOperator) bool {
	if value < min || value > max {
		return false
	}
	return true
}

// 根据 输入的 int类型值，判断 数据是否在 [min max] 之间
func (c Compare) BetweenInt(value, min, max int, op ComparisonOperator) bool {
	if value < min || value > max {
		return false
	}
	return true
}

// 判断 value是否在target之后
func (c Compare) CompareIsAfterTime(value, target time.Time) bool {
	return value.After(target)
}

// 判断 value是否在target之前
func (c Compare) CompareIsBeforeTime(value, target time.Time) bool {
	return value.Before(target)
}
