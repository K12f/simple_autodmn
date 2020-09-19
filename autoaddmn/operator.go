package autoaddmn

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
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
	Mul ArithmeticOperator = "*"
	Div ArithmeticOperator = "/"
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

	if len(value) != 2 {
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

func (o Operator) MultiLogicCombine(compared []bool, logicBox []LogicOperatorType) (bool, error) {
	var err error
	var flag bool

	if len(compared) == 1 {
		return compared[0], err
	}
	// 防止溢出
	if len(logicBox) > 1 && len(logicBox) >= len(compared) {
		logicBox = logicBox[:len(logicBox)-1]
	}

	for k, c := range logicBox {
		if len(compared) < k {
			return false, ParseRuleOutOfRangeErr
		}

		flag, err = o.LogicCombine(c, compared[k], compared[k+1])
		if err != nil {
			return false, errors.WithStack(err)
		}
		compared[k+1] = flag
	}
	return flag, err
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

// 乘法
func (o Operator) MulInt(a, b int) int {
	return a * b
}

func (o Operator) MulFloat64(a, b float64) float64 {
	return a * b
}

func (o Operator) DivFloat64(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	value := a / b
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// 百分比
func (o Operator) Percent(a interface{}, per float64) float64 {
	tmp := a.(float64)
	return tmp * per
}

// 增加时间,以分钟为单位
func (o Operator) AddTime(a time.Time, minutes float64) time.Time {
	b := time.Duration(minutes) * time.Minute
	return a.Add(b)
}

// 减少时间,以分钟为单位
func (o Operator) SubTime(a time.Time, minutes float64) time.Time {
	b := time.Duration(minutes) * time.Minute
	return a.Add(-b)
}
