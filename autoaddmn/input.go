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
func (i Input) GetInt(value interface{}) (int, error) {
	var ret int
	ret, ok := value.(int)
	if ok {
		return ret, nil
	}
	tmp, ok := value.(string)
	if !ok {
		return 0, errors.WithStack(CouldNotParseInputErr)
	}
	ret, err := strconv.Atoi(tmp)
	if err != nil {
		return 0, errors.Wrap(err, CouldNotParseInputErr.Error())
	}
	return ret, nil
}

// 获取 float64类型的输入
func (i Input) GetFloat64(value interface{}) (float64, error) {
	var ret float64
	ret, ok := value.(float64)
	if ok {
		return ret, nil
	}
	valueT, ok := value.(int)
	if ok {
		ret = float64(valueT)
		return ret, nil
	}
	tmp, ok := value.(string)
	if !ok {
		return 0, errors.WithStack(CouldNotParseInputErr)
	}
	ret, err := strconv.ParseFloat(tmp, 64)
	if err != nil {
		return 0, errors.Wrap(err, CouldNotParseInputErr.Error())
	}
	return ret, nil
}

// 获取字符串类型的输入
func (i Input) GetString(value interface{}) (string, error) {
	ret, ok := value.(string)
	if !ok {
		return "", errors.WithStack(CouldNotParseInputErr)
	}
	return ret, nil
}
