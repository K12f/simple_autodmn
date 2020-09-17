package autoaddmn

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	ParamsErr = errors.New("参数有误")

	//operator
	// 逻辑运算符
	UnknownLogicOperatorErr      = errors.New("未知的逻辑运算符")
	UnknownArithmeticOperatorErr = errors.New("未知的算数运算符")
	UnknownValueTypeErr          = errors.New("未知的值类型")

	// 解析 rule
	NotFoundRuleComponentsErr      = errors.New("未发现广告规则组件数据")
	NotFoundRuleComponentsLeftErr  = errors.New("未发现广告规则组件左节点数据")
	NotFoundRuleComponentsRightErr = errors.New("未发现广告规则组件右节点数据")
	ParseRuleOutOfRangeErr         = errors.New("解析规则超过范围")
	CouldNotParseRulesErr          = errors.New("不能解析规则")
	CouldNotParseDecisionsErr      = errors.New("不能解析决策")
	CouldNotParseOrdersErr         = errors.New("不能解析指令")

	CouldNotFindParseRulesErr     = errors.New("没有发现解析规则")
	CouldNotFindParseDecisionsErr = errors.New("没有发现解析决策")
	CouldNotFindParseOrdersErr    = errors.New("没有发现解析指令")
	//compare

	CouldNotCompareDateNotParamsErr = fmt.Errorf("不能比较时间:%w", ParamsErr)
	CouldNotCompareParseErr         = errors.New("解析失败")
	CouldNotParseInputErr           = errors.New("解析失败:不能获取到input值")

	//value
	CouldNotParseValueErr = errors.New("不能解析值类型")
)
