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

func (c Compare) EqFloat64(a, b float64) bool {
	return a == b
}

func (c Compare) GTFloat64(a, b float64) bool {
	return a > b
}

func (c Compare) GETFloat64(a, b float64) bool {
	return a >= b
}

func (c Compare) LTFloat64(a, b float64) bool {
	return a < b
}

func (c Compare) LETFloat64(a, b float64) bool {
	return a <= b
}

func (c Compare) BetweenFloat64(a float64, b ...float64) bool {
	return a >= b[0] && a <= b[1]
}

func (c Compare) EqDate(a, b time.Time) bool {
	return a.Equal(b)
}

func (c Compare) LTDate(a, b time.Time) bool {
	return a.Unix() < b.Unix()
}

func (c Compare) GTDate(a, b time.Time) bool {
	return a.Unix() > b.Unix()
}

func (c Compare) LETDate(a, b time.Time) bool {
	return a.Before(b)
}

func (c Compare) GETDate(a, b time.Time) bool {
	return a.After(b)
}

func (c Compare) BetweenDate(a, b, m time.Time) bool {
	return a.After(b) && a.Before(m)
}
