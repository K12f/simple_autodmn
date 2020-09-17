package autoaddmn

import (
	"github.com/pkg/errors"
	"strconv"
)

// 输入
type Input struct {
	Value interface{} `json:"value"` // 输入值
}

func NewInput() *Input {
	return &Input{}
}

// 获取int类型的输入值
func (i Input) GetInt(input Input) (int, error) {
	var value int
	value, ok := input.Value.(int)
	if ok {
		return value, nil
	}
	tmp, ok := input.Value.(string)
	if !ok {
		return 0, errors.WithStack(CouldNotParseInputErr)
	}
	value, err := strconv.Atoi(tmp)
	if err != nil {
		return 0, errors.Wrap(err, CouldNotParseInputErr.Error())
	}
	return value, nil
}

// 获取 float64类型的输入
func (i Input) GetFloat64(input Input) (float64, error) {
	var value float64
	value, ok := input.Value.(float64)
	if ok {
		return value, nil
	}
	valueT, ok := input.Value.(int)
	if ok {
		value = float64(valueT)
		return value, nil
	}
	tmp, ok := input.Value.(string)
	if !ok {
		return 0, errors.WithStack(CouldNotParseInputErr)
	}
	value, err := strconv.ParseFloat(tmp, 64)
	if err != nil {
		return 0, errors.Wrap(err, CouldNotParseInputErr.Error())
	}
	return value, nil
}

// 获取字符串类型的输入
func (i Input) GetString(input Input) (string, error) {
	value, ok := input.Value.(string)
	if !ok {
		return "", errors.WithStack(CouldNotParseInputErr)
	}
	return value, nil
}
