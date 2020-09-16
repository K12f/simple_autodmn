package autodmn

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

var (
	operator = NewOperator()
	compare  = NewCompare()
	inp      = NewInput()
)

const TimeFormat = "2006-01-02 15:04:05"

type ValueType string

const (
	OI  ValueType = "int"
	OF  ValueType = "float64"
	OD  ValueType = "date"
	OMD ValueType = "mul_date"
)

type Value struct {
	Name      string      `json:"name"`       // 名称
	Desc      string      `json:"desc"`       // 描述
	Slot      string      `json:"slot"`       // 变量标志
	Value     interface{} `json:"value"`      // 值
	ValueType ValueType   `json:"value_type"` //值的类型
}

func NewValue() *Value {
	return &Value{}
}

// 规则有效时间
type AdRuleEnableTime struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// 继续投放时长
type AdKeepPutTime struct {
	Value float64 `json:"value"` // 小时
}

// 广告起投时长
type AdLeastTime struct {
	Value float64 `json:"value"` // 小时
}

// 广告消耗速度
type AdSpeedRate struct {
	Value int `json:"value"` //广告消耗速度
}

// 广告转化成本
type AdConvCost struct {
	Value float64 `json:"value"` //分
}

// 广告当前出价
type AdCurCost struct {
	Value float64 `json:"value"` //分
}

// 广告曝光速度
type AdExpoSpeed struct {
	Value int `json:"value"` //
}

// 广告当日花费
type AdDayCost struct {
	Value float64 `json:"value"` //分
}

// 广告目标转化率
type ACR struct {
	Value float64 `json:"value"` // 百分比
}

// 账户可用余额
type Account struct {
	Value float64 `json:"value"` //分
}

func (v *Value) Compare(inputs []Input, op ComparisonOperator) (bool, error) {
	switch v.ValueType {
	case OI:
		return v.compareInt(inputs, op)
	case OF:
		return v.compareFloat64(inputs, op)
	case OD:
		return v.compareDate(inputs, op)
	case OMD:
		return v.compareDate(inputs, op)
	default:
		return false, errors.WithStack(CouldNotParseValueErr)
	}
}

func (v *Value) compareDate(inputs []Input, op ComparisonOperator) (bool, error) {
	var flag bool
	if len(inputs) != 2 {
		return false, CouldNotCompareDateNotParamsErr
	}
	inpStart, err := inp.GetString(inputs[0])
	if err != nil {
		return false, err
	}
	formattedStart, err := time.Parse(TimeFormat, inpStart)
	if err != nil {
		return false, errors.WithStack(fmt.Errorf("input格式化开始时间%w", err))
	}
	inpEnd, err := inp.GetString(inputs[1])
	if err != nil {
		return false, err
	}
	formattedEnd, err := time.Parse(TimeFormat, inpEnd)
	if err != nil {
		return false, errors.WithStack(fmt.Errorf("input格式化结束时间%w", err))
	}

	enableTime, err := v.GetValue()
	if err != nil {
		return false, errors.WithStack(err)
	}
	if len(enableTime) < 1 {
		return false, errors.WithStack(CouldNotCompareParseErr)
	}

	enableTimeStart, ok := enableTime[0].(string)
	if !ok {
		return false, errors.WithStack(CouldNotCompareParseErr)
	}
	formattedEnableStart, err := time.Parse(TimeFormat, enableTimeStart)
	if err != nil {
		return false, errors.WithStack(CouldNotCompareParseErr)
	}
	if len(enableTime) == 2 {
		enableTimeEnd, ok := enableTime[1].(string)
		if !ok {
			return false, errors.WithStack(CouldNotCompareParseErr)
		}

		formattedEnableEnd, err := time.Parse(TimeFormat, enableTimeEnd)
		if err != nil {
			return false, errors.WithStack(CouldNotCompareParseErr)
		}
		if compare.CompareIsBeforeTime(formattedEnableEnd, formattedEnableStart) {
			return false, errors.New("结束时间不能小于开始时间")
		}
		if compare.CompareIsBeforeTime(formattedEnd, formattedStart) {
			return false, errors.New("结束时间不能小于开始时间")
		}

		flag = compare.CompareIsAfterTime(formattedEnableStart, formattedStart) && compare.CompareIsAfterTime(formattedEnd, formattedEnableEnd)
	} else {
		flag = compare.CompareIsAfterTime(formattedEnableStart, formattedStart) && compare.CompareIsBeforeTime(formattedEnableStart, formattedEnd)
	}

	return flag, nil
}

func (v *Value) compareFloat64(inputs []Input, op ComparisonOperator) (bool, error) {
	if len(inputs) > 2 || len(inputs) < 1 {
		return false, ParamsErr
	}

	value, err := v.GetValue()
	if err != nil {
		return false, err
	}
	valueTmp, ok := value[0].(float64)
	if !ok {
		return false, errors.WithStack(CouldNotParseValueErr)
	}

	input, err := inp.GetFloat64(inputs[0])
	if err != nil {
		return false, errors.WithStack(err)
	}
	if len(inputs) == 1 {
		return compare.CompareFloat64(valueTmp, input, op), nil
	} else {
		input2, err := inp.GetFloat64(inputs[1])
		if err != nil {
			return false, errors.WithStack(err)
		}
		return compare.BetweenFloat64(valueTmp, input, input2, op), nil
	}
}

func (v *Value) compareInt(inputs []Input, op ComparisonOperator) (bool, error) {
	if len(inputs) > 2 || len(inputs) < 1 {
		return false, errors.WithStack(ParamsErr)
	}

	value, err := v.GetValue()
	if err != nil {
		return false, err
	}
	valueTmp, ok := value[0].(int)
	if !ok {
		return false, errors.WithStack(CouldNotParseValueErr)
	}

	input, err := inp.GetInt(inputs[0])
	if err != nil {
		return false, errors.WithStack(err)
	}

	if len(inputs) == 1 {
		return compare.CompareInt(valueTmp, input, op), nil
	} else {
		input2, err := inp.GetInt(inputs[1])
		if err != nil {
			return false, errors.WithStack(err)
		}
		return compare.BetweenInt(valueTmp, input, input2, op), nil
	}
}

func (v *Value) Calc(inputs []Input, op ArithmeticOperator) ([]interface{}, error) {
	var result []interface{}
	switch v.ValueType {
	case OI:
		return v.calcInt(inputs, op)
	case OF:
		return v.calcFloat64(inputs, op)
	case OMD:
		return v.calcTime(inputs, op)
	default:
		return result, errors.WithStack(UnknownValueTypeErr)
	}
}

func (v *Value) calcTime(inputs []Input, op ArithmeticOperator) ([]interface{}, error) {
	var result []interface{}
	if len(inputs) != 1 {
		return result, ParamsErr
	}
	// 增加的以分钟为单位
	input := time.Duration(inputs[0].Value.(int))

	value, err := v.GetValue()
	if err != nil {
		return result, err
	}

	if len(value) != 2 {
		return result, ParamsErr
	}

	start := value[0].(string)
	end := value[1].(string)

	switch op {
	case ADD:
		tmpStart, err := time.Parse(TimeFormat, start)
		if err != nil {
			return result, errors.Wrap(err, err.Error())
		}
		start = tmpStart.Add(time.Minute * input).Format(TimeFormat)

		tmpEnd, err := time.Parse(TimeFormat, end)
		if err != nil {
			return result, errors.Wrap(err, err.Error())
		}
		end = tmpEnd.Add(time.Minute * input).Format(TimeFormat)
	case SUB:
		tmpStart, err := time.Parse(TimeFormat, start)
		if err != nil {
			return result, errors.Wrap(err, err.Error())
		}
		start = tmpStart.Add(-time.Minute * input).Format(TimeFormat)

		tmpEnd, err := time.Parse(TimeFormat, end)
		if err != nil {
			return result, errors.Wrap(err, err.Error())
		}
		end = tmpEnd.Add(-time.Minute * input).Format(TimeFormat)
	default:
		return result, errors.WithStack(UnknownArithmeticOperatorErr)
	}
	result = append(result, start)
	result = append(result, end)
	return result, nil
}

func (v *Value) calcFloat64(inputs []Input, op ArithmeticOperator) ([]interface{}, error) {
	var result []interface{}
	if len(inputs) != 1 {
		return result, ParamsErr
	}
	input, err := inp.GetFloat64(inputs[0])
	if err != nil {
		return result, errors.WithStack(err)
	}

	value, err := v.GetValue()
	if err != nil {
		return result, errors.WithStack(err)
	}
	valueTmp, ok := value[0].(float64)
	if !ok {
		return result, errors.WithStack(CouldNotParseValueErr)
	}

	switch op {
	case ADD:
		result = append(result, operator.AddFloat64(valueTmp, input))
	case SUB:
		result = append(result, operator.SubFloat64(valueTmp, input))
	case Per:
		result = append(result, operator.Percent(valueTmp, input))
	default:
		return result, errors.WithStack(UnknownArithmeticOperatorErr)
	}
	return result, nil
}

func (v *Value) calcInt(inputs []Input, op ArithmeticOperator) ([]interface{}, error) {
	var result []interface{}
	if len(inputs) != 1 {
		return result, errors.WithStack(ParamsErr)
	}

	value, err := v.GetValue()
	if err != nil {
		return result, errors.WithStack(err)
	}
	valueTmp, ok := value[0].(int)
	if !ok {
		return result, errors.WithStack(CouldNotParseValueErr)
	}
	switch op {
	case ADD:
		input, err := inp.GetInt(inputs[0])
		if err != nil {
			return result, errors.WithStack(err)
		}
		result = append(result, operator.AddInt(valueTmp, input))
	case SUB:
		input, err := inp.GetInt(inputs[0])
		if err != nil {
			return result, errors.WithStack(err)
		}
		result = append(result, operator.SubInt(valueTmp, input))
	case Per:
		input, err := inp.GetFloat64(inputs[0])
		if err != nil {
			return result, errors.WithStack(err)
		}
		result = append(result, operator.Percent(valueTmp, input))
	default:
		return result, errors.WithStack(UnknownArithmeticOperatorErr)
	}
	return result, nil
}

func (v *Value) SetValue(value []interface{}) (interface{}, error) {
	var result interface{}
	switch t := v.Value.(type) {
	case AdRuleEnableTime:
		result = AdRuleEnableTime{
			Start: value[0].(string),
			End:   value[1].(string),
		}
	case AdKeepPutTime:
		result = AdKeepPutTime{
			Value: value[0].(float64),
		}
	case AdLeastTime:
		result = AdLeastTime{
			Value: value[0].(float64),
		}
	case AdConvCost:
		result = AdConvCost{
			Value: value[0].(float64),
		}
	case AdCurCost:
		result = AdCurCost{
			Value: value[0].(float64),
		}
	case AdDayCost:
		result = AdDayCost{
			Value: value[0].(float64),
		}
	case ACR:
		result = ACR{
			Value: value[0].(float64),
		}
	case Account:
		result = Account{
			Value: value[0].(float64),
		}
	case AdSpeedRate:
		result = AdSpeedRate{
			Value: value[0].(int),
		}
	case AdExpoSpeed:
		result = AdExpoSpeed{
			Value: value[0].(int),
		}
	default:
		return result, errors.WithStack(fmt.Errorf("未知的类型:%t", t))
	}
	return result, nil
}

func (v *Value) GetValue() ([]interface{}, error) {
	var value []interface{}
	var err error
	switch t := v.Value.(type) {
	case AdRuleEnableTime:
		value = append(value, v.Value.(AdRuleEnableTime).Start)
		value = append(value, v.Value.(AdRuleEnableTime).End)
	case AdKeepPutTime:
		value = append(value, v.Value.(AdKeepPutTime).Value)
	case AdLeastTime:
		value = append(value, v.Value.(AdLeastTime).Value)
	case AdConvCost:
		value = append(value, v.Value.(AdConvCost).Value)
	case AdCurCost:
		value = append(value, v.Value.(AdCurCost).Value)
	case AdDayCost:
		value = append(value, v.Value.(AdDayCost).Value)
	case ACR:
		value = append(value, v.Value.(ACR).Value)
	case Account:
		value = append(value, v.Value.(Account).Value)
	case AdSpeedRate:
		value = append(value, v.Value.(AdSpeedRate).Value)
	case AdExpoSpeed:
		value = append(value, v.Value.(AdExpoSpeed).Value)
	default:
		err = fmt.Errorf("未知的类型:%t", t)
	}
	return value, errors.WithStack(err)
}
