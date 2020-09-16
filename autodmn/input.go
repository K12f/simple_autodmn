package autodmn

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
		return 0, errors.WithStack(err)
	}
	return value, nil
}

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
		return 0, errors.WithStack(err)
	}
	return value, nil
}

func (i Input) GetString(input Input) (string, error) {
	value, ok := input.Value.(string)
	if !ok {
		return "", errors.WithStack(CouldNotParseInputErr)
	}
	return value, nil
}
