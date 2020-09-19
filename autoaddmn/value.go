package autoaddmn

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

var (
	inp      = NewInput()
	compare  = NewCompare()
	operator = NewOperator()
)

// 值类型
type ValueType string

const (
	//OI  ValueType = "int"
	OF  ValueType = "float64"
	OD  ValueType = "date"
	OMD ValueType = "mul_date"
)

// 值结构体
type Value struct {
	Slot      AdInfoSlot  `json:"slot"`       // 变量唯一标记, 根据此标记去匹配对应的 AdInfo中的数据
	Value     interface{} `json:"value"`      // 值
	ValueType ValueType   `json:"value_type"` // 值的类型   ,用来匹配对的类型，减少适配
}

// 根据输入比较值
func (v *Value) Compare(value []interface{}, valueType ValueType, inputs []Input, op ComparisonOperator) (bool, error) {
	switch valueType {
	//case OI:
	//    // int
	//    return v.compareFloat64(value, inputs, op)
	case OF:
		// float
		return v.compareFloat64(value, inputs, op)
	case OD:
		// 日期
		return v.compareDate(value, inputs, op)
	case OMD:
		// 多个日期
		return v.compareDate(value, inputs, op)
	default:
		return false, errors.WithStack(CouldNotParseValueErr)
	}
}

// 比较日期
func (v *Value) compareDate(value []interface{}, inputs []Input, op ComparisonOperator) (bool, error) {
	var flag bool
	if len(inputs) != 2 {
		return false, CouldNotCompareDateNotParamsErr
	}
	inpStart, err := inp.GetString(inputs[0].Value)
	if err != nil {
		return false, err
	}
	formattedStart, err := time.Parse(TimeFormat, inpStart)
	if err != nil {
		return false, errors.WithStack(fmt.Errorf("input格式化开始时间%w", err))
	}
	inpEnd, err := inp.GetString(inputs[1].Value)
	if err != nil {
		return false, err
	}
	formattedEnd, err := time.Parse(TimeFormat, inpEnd)
	if err != nil {
		return false, errors.WithStack(fmt.Errorf("input格式化结束时间%w", err))
	}

	if len(value) < 1 {
		return false, errors.WithStack(CouldNotCompareParseErr)
	}

	enableTimeStart, err := inp.GetString(value[0])
	if err != nil {
		return false, err
	}
	formattedEnableStart, err := time.Parse(TimeFormat, enableTimeStart)
	if err != nil {
		return false, errors.WithStack(CouldNotCompareParseErr)
	}

	if compare.CompareIsBeforeTime(formattedEnd, formattedStart) {
		return false, errors.New("结束时间不能小于开始时间")
	}

	// 如果 注入的值是两个
	if len(value) == 2 {
		enableTimeEnd, err := inp.GetString(value[1])
		if err != nil {
			return false, err
		}

		formattedEnableEnd, err := time.Parse(TimeFormat, enableTimeEnd)
		if err != nil {
			return false, errors.WithStack(CouldNotCompareParseErr)
		}
		if compare.CompareIsBeforeTime(formattedEnableEnd, formattedEnableStart) {
			return false, errors.New("结束时间不能小于开始时间")
		}

		flag = compare.CompareIsAfterTime(formattedEnableStart, formattedStart) && compare.CompareIsAfterTime(formattedEnd, formattedEnableEnd)
	} else {
		flag = compare.CompareIsAfterTime(formattedEnableStart, formattedStart) && compare.CompareIsBeforeTime(formattedEnableStart, formattedEnd)
	}

	return flag, nil
}

// 比较float64类型
func (v *Value) compareFloat64(value []interface{}, inputs []Input, op ComparisonOperator) (bool, error) {
	if len(inputs) > 2 || len(inputs) < 1 {
		return false, ParamsErr
	}

	valueTmp, err := inp.GetFloat64(value[0])
	if err != nil {
		return false, errors.Wrap(err, CouldNotParseValueErr.Error())
	}

	input, err := inp.GetFloat64(inputs[0].Value)
	if err != nil {
		return false, errors.WithStack(err)
	}
	if len(inputs) == 1 {
		return compare.CompareFloat64(valueTmp, input, op), nil
	} else {
		// 多输入个值使用between 判断
		input2, err := inp.GetFloat64(inputs[1].Value)
		if err != nil {
			return false, errors.WithStack(err)
		}
		return compare.BetweenFloat64(valueTmp, input, input2, op), nil
	}
}

// 根据操作符和输入值 计算值，并返回
func (v *Value) Calc(value []interface{}, valueType ValueType, inputs []Input, op ArithmeticOperator) ([]interface{}, error) {
	var result []interface{}

	switch valueType {
	//case OI:
	//    return v.calcFloat64(value, inputs, op)
	case OF:
		return v.calcFloat64(value, inputs, op)
	case OD:
		return v.calcTime(value, inputs, op)
	case OMD:
		return v.calcTime(value, inputs, op)
	default:
		return result, errors.WithStack(UnknownValueTypeErr)
	}
}

// 计算时间
func (v *Value) calcTime(value []interface{}, inputs []Input, op ArithmeticOperator) ([]interface{}, error) {
	var result []interface{}
	if len(inputs) != 1 {
		return result, ParamsErr
	}
	// 增加的以分钟为单位
	input, err := inp.GetInt(inputs[0].Value)
	if err != nil {
		return result, errors.WithStack(err)
	}

	if len(value) != 2 {
		return result, ParamsErr
	}

	start, err := inp.GetString(value[0])
	if err != nil {
		return result, errors.WithStack(err)
	}
	end, err := inp.GetString(value[1])
	if err != nil {
		return result, errors.WithStack(err)
	}
	switch op {
	case ADD:
		tmpStart, err := time.Parse(TimeFormat, start)
		if err != nil {
			return result, errors.Wrap(err, err.Error())
		}
		start = operator.AddTime(tmpStart, input).Format(TimeFormat)
		tmpEnd, err := time.Parse(TimeFormat, end)
		if err != nil {
			return result, errors.Wrap(err, err.Error())
		}
		end = operator.AddTime(tmpEnd, input).Format(TimeFormat)
	case SUB:
		tmpStart, err := time.Parse(TimeFormat, start)
		if err != nil {
			return result, errors.Wrap(err, err.Error())
		}
		start = operator.SubTime(tmpStart, input).Format(TimeFormat)
		tmpEnd, err := time.Parse(TimeFormat, end)
		if err != nil {
			return result, errors.Wrap(err, err.Error())
		}
		end = operator.SubTime(tmpEnd, input).Format(TimeFormat)
	default:
		return result, errors.WithStack(UnknownArithmeticOperatorErr)
	}
	result = append(result, start)
	result = append(result, end)
	return result, nil
}

// 计算float
func (v *Value) calcFloat64(value []interface{}, inputs []Input, op ArithmeticOperator) ([]interface{}, error) {
	var result []interface{}
	if len(inputs) != 1 {
		return result, ParamsErr
	}
	input, err := inp.GetFloat64(inputs[0].Value)
	if err != nil {
		return result, errors.WithStack(err)
	}

	valueTmp, err := inp.GetFloat64(value[0])

	if err != nil {
		return result, errors.WithStack(err)
	}

	switch op {
	case ADD:
		result = append(result, operator.AddFloat64(valueTmp, input))
	case SUB:
		result = append(result, operator.SubFloat64(valueTmp, input))
	case Mul:
		result = append(result, operator.MulFloat64(valueTmp, input))
	case Div:
		result = append(result, operator.DivFloat64(valueTmp, input))
	case Per:
		result = append(result, operator.Percent(valueTmp, input))
	default:
		return result, errors.WithStack(UnknownArithmeticOperatorErr)
	}
	return result, nil
}
