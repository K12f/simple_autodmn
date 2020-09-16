package autodmn

import (
	"time"
)

type Operator struct {
}

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

// 条件类型
type LogicOperatorType string

const (
	AND LogicOperatorType = "&&" //条件操作符   condition
	OR  LogicOperatorType = "||" //条件操作符   condition
	NOT LogicOperatorType = "!"  //条件操作符   condition
)

type LogicOperator struct {
}

type ArithmeticOperator string

const (
	ADD ArithmeticOperator = "+"
	SUB ArithmeticOperator = "-"
	Per ArithmeticOperator = "%"
)

func NewOperator() *Operator {
	return &Operator{}
}

func (o Operator) LogicCombine(op LogicOperatorType, value ...bool) (bool, error) {

	if len(value) > 2 || len(value) < 1 {
		return false, ParamsErr
	}
	switch op {
	case AND:
		return value[0] && value[1], nil
	case OR:
		return value[0] || value[1], nil
	case NOT:
		return !value[0], nil
	default:
		return false, UnknownLogicOperatorErr
	}
}

//func (o Operator) CalcInt(a, b interface{}, op ArithmeticOperator) (interface{}, error) {
//    tmpA := a.(float64)
//    tmpB := b.(float64)
//
//    switch op {
//    case ADD:
//        return o.AddFloat64(tmpA, tmpB), nil
//    case SUB:
//        return o.SubFloat64(tmpA, tmpB), nil
//    case Per:
//        return o.Percent(tmpA, tmpB), nil
//    case :
//
//    default:
//        return nil, errors.New("未知的运算符类型")
//    }
//}

func (o Operator) AddInt(a, b int) int {
	return a + b
}

func (o Operator) AddFloat64(a, b float64) float64 {
	return a + b
}

func (o Operator) SubInt(a, b int) int {
	return a - b
}

func (o Operator) SubFloat64(a, b float64) float64 {
	return a - b
}

func (o Operator) Percent(a interface{}, per float64) float64 {
	tmp := a.(float64)
	return tmp * per
}

func (o Operator) AddTime(a time.Time, b time.Duration) time.Time {
	return a.Add(b * time.Minute)
}

func (o Operator) SubTime(a time.Time, b time.Duration) time.Time {
	return a.Add(-b)
}
