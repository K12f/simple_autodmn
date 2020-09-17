package autoaddmn

import (
	"time"
)

// 逻辑操作符类型
type LogicOperatorType string

// 逻辑操作符
const (
	// 与
	AND LogicOperatorType = "&&"
	// 或
	OR LogicOperatorType = "||"
	// 非
	//NOT LogicOperatorType = "!"
)

// 算数运算符类型
type ArithmeticOperator string

const (
	ADD ArithmeticOperator = "+"
	SUB ArithmeticOperator = "-"
	Per ArithmeticOperator = "%"
)

// 操作符
type Operator struct {
}

func NewOperator() *Operator {
	return &Operator{}
}

// 根据 && || ,合并bool类型的值，返回最终值
func (o Operator) LogicCombine(op LogicOperatorType, value ...bool) (bool, error) {

	if len(value) > 2 || len(value) < 1 {
		return false, ParamsErr
	}
	switch op {
	case AND:
		return value[0] && value[1], nil
	case OR:
		return value[0] || value[1], nil
	//case NOT:
	//	return !value[0], nil
	default:
		return false, UnknownLogicOperatorErr
	}
}

// 加
func (o Operator) AddInt(a, b int) int {
	return a + b
}

func (o Operator) AddFloat64(a, b float64) float64 {
	return a + b
}

// 减
func (o Operator) SubInt(a, b int) int {
	return a - b
}

func (o Operator) SubFloat64(a, b float64) float64 {
	return a - b
}

// 百分比
func (o Operator) Percent(a interface{}, per float64) float64 {
	tmp := a.(float64)
	return tmp * per
}

// 增加时间,以分钟为单位
func (o Operator) AddTime(a time.Time, minutes int) time.Time {
	b := time.Duration(minutes) * time.Minute
	return a.Add(b)
}

// 减少时间,以分钟为单位
func (o Operator) SubTime(a time.Time, minutes int) time.Time {
	b := time.Duration(minutes) * time.Minute
	return a.Add(-b)
}
