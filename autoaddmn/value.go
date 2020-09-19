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
	var resultTemp bool
	var compared []bool
	var logicBox []LogicOperatorType
	var err error
	var result bool

	// 不能没有值和输入 且 输入的长度>=值的长度
	if len(inputs) == 0 || len(value) == 0 || (len(inputs) < len(value) && len(inputs) != 1) {
		return result, ParamsErr
	}

	// 比较的输入值 长度不能超过2
	if len(inputs) > 2 {
		return result, errors.WithStack(CouldOutOfRangeFor2)
	}

	for _, vV := range value {
		// 值不能为空
		if vV == nil {
			return result, errors.WithStack(CouldNotFindValueErr)
		}

		switch valueType {
		//case OI:
		//    // int
		//    return v.compareFloat64(value, inputs, op)
		case OF:
			// float
			resultTemp, err = v.compareFloat64(vV, inputs, op)
		case OD, OMD:
			// 日期
			resultTemp, err = v.compareDate(vV, inputs, op)
		default:
			err = errors.WithStack(CouldNotParseValueErr)
		}
		compared = append(compared, resultTemp)
		logicBox = append(logicBox, AND)
	}
	result, err = operator.MultiLogicCombine(compared, logicBox)
	return result, err

}

func (v *Value) compareDate(value interface{}, inputs []Input, op ComparisonOperator) (bool, error) {
	var input []string
	var err error
	var result bool

	valueTmp, err := inp.GetString(value)
	if err != nil {
		return false, errors.Wrap(err, CouldNotParseValueErr.Error())
	}

	for _, inpValue := range inputs {
		inputTemp, err := inp.GetString(inpValue.Value)
		if err != nil {
			return false, errors.WithStack(err)
		}
		input = append(input, inputTemp)
	}

	//now, _ := time.Parse(format, time.Now().Format(format))
	timeA, err := time.Parse(TimeFormat, valueTmp)
	if err != nil {
		return false, errors.WithStack(fmt.Errorf("value格式化时间%w", err))
	}

	timeB, err := time.Parse(TimeFormat, input[0])
	if err != nil {
		return false, errors.WithStack(fmt.Errorf("input格式化时间%w", err))
	}

	switch op {
	case EQ:
		result = compare.EqDate(timeA, timeB)
	case LT:
		result = compare.LTDate(timeA, timeB)
	case GT:
		result = compare.GTDate(timeA, timeB)
	case LET:
		result = compare.LETDate(timeA, timeB)
	case GET:
		result = compare.GETDate(timeA, timeB)
	case BETWEEN:
		timeC, err := time.Parse(TimeFormat, input[1])
		if err != nil {
			return false, errors.WithStack(fmt.Errorf("input格式化时间%w", err))
		}
		result = compare.BetweenDate(timeA, timeB, timeC)
	default:

	}
	return result, err
}

// 比较float64类型
func (v *Value) compareFloat64(value interface{}, inputs []Input, op ComparisonOperator) (bool, error) {
	var input []float64
	var err error
	var result bool

	valueTmp, err := inp.GetFloat64(value)
	if err != nil {
		return result, errors.Wrap(err, CouldNotParseValueErr.Error())
	}

	for _, inpValue := range inputs {
		inputTemp, err := inp.GetFloat64(inpValue.Value)
		if err != nil {
			return result, errors.WithStack(err)
		}
		input = append(input, inputTemp)
	}

	switch op {
	case EQ:
		result = compare.EqFloat64(valueTmp, input[0])
	case LT:
		result = compare.LTFloat64(valueTmp, input[0])
	case GT:
		result = compare.GTFloat64(valueTmp, input[0])
	case LET:
		result = compare.LETFloat64(valueTmp, input[0])
	case GET:
		result = compare.GETFloat64(valueTmp, input[0])
	case BETWEEN:
		result = compare.BetweenFloat64(valueTmp, input[0])
	default:

	}
	return result, err
}

// 根据操作符和输入值 计算值，并返回
func (v *Value) Calc(value []interface{}, valueType ValueType, inputs []Input, op ArithmeticOperator) ([]interface{}, error) {
	var result []interface{}
	var temp interface{}
	var input Input
	var err error

	// 不能没有值和输入 且 输入的长度>=值的长度
	if len(inputs) == 0 || len(value) == 0 || (len(inputs) < len(value) && len(inputs) != 1) {
		return result, ParamsErr
	}

	for k, vV := range value {
		// 值不能为空
		if vV == nil {
			return result, errors.WithStack(CouldNotFindValueErr)
		}
		//如果只有一个值，那么就获取该值，否则获取对应的 输入值
		if len(inputs) == 1 {
			// 增加的以分钟为单位
			input = inputs[0]
		} else {
			input = inputs[k]
		}

		switch valueType {
		//case OI:
		//    return v.calcFloat64(value, inputs, op)
		case OF:
			temp, err = v.calcFloat64(vV, input, op)
			if err != nil {
				return result, errors.WithStack(err)
			}
		case OD:
			temp, err = v.calcTime(vV, input, op)
			if err != nil {
				return result, errors.WithStack(err)
			}
		case OMD:
			temp, err = v.calcTime(vV, input, op)
			if err != nil {
				return result, errors.WithStack(err)
			}
		default:
			return result, errors.WithStack(UnknownValueTypeErr)
		}
		result = append(result, temp)
	}

	return result, nil
}

// 计算时间
func (v *Value) calcTime(value interface{}, input Input, op ArithmeticOperator) (string, error) {
	var result string
	var timeTemp time.Time

	temp, err := inp.GetString(value)
	if err != nil {
		return result, errors.WithStack(err)
	}

	inputTemp, err := inp.GetFloat64(input.Value)
	if err != nil {
		return result, errors.WithStack(err)
	}

	timeTemp, err = time.Parse(TimeFormat, temp)
	if err != nil {
		return result, errors.Wrap(err, err.Error())
	}

	switch op {
	case ADD:
		result = operator.AddTime(timeTemp, inputTemp).Format(TimeFormat)
	case SUB:
		result = operator.SubTime(timeTemp, inputTemp).Format(TimeFormat)
	default:
		return result, errors.WithStack(UnknownArithmeticOperatorErr)
	}
	return result, nil
}

// 计算float
func (v *Value) calcFloat64(value interface{}, input Input, op ArithmeticOperator) (float64, error) {
	var result float64
	valueTmp, err := inp.GetFloat64(value)

	if err != nil {
		return result, errors.WithStack(err)
	}

	inputTemp, err := inp.GetFloat64(input.Value)
	if err != nil {
		return result, errors.WithStack(err)
	}

	switch op {
	case ADD:
		result = operator.AddFloat64(valueTmp, inputTemp)
	case SUB:
		result = operator.SubFloat64(valueTmp, inputTemp)
	case Mul:
		result = operator.MulFloat64(valueTmp, inputTemp)
	case Div:
		result = operator.DivFloat64(valueTmp, inputTemp)
	case Per:
		result = operator.Percent(valueTmp, inputTemp)
	default:
		return result, errors.WithStack(UnknownArithmeticOperatorErr)
	}
	return result, nil
}
