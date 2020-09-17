package autoaddmn

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

// 解析器，用来解析规则，决策，指令
type Parse struct {
}

func NewParse() *Parse {
	return &Parse{}
}

// 根据规则，广告微信数据信息，解析规则
func (p *Parse) ParseRules(rules []*Rule, ad *Ad) (bool, error) {
	if len(rules) < 1 {
		return false, errors.WithStack(CouldNotFindParseRulesErr)
	}
	fmt.Println("正在解析规则")
	// 结果集
	var compared []bool
	// 操作符
	var logicBox []LogicOperatorType
	var err error
	var flag bool
	operator := NewOperator()

	for _, rv := range rules {
		//获取
		value, err := ad.AdInfo.GetValue(rv.Value.Slot)
		if err != nil {
			return false, errors.WithStack(err)
		}
		// 比较值
		rvCompared, err := rv.Value.Compare(value, rv.Value.ValueType, rv.Inputs, rv.COperator)

		if err != nil {
			return false, errors.WithStack(err)
		}
		compared = append(compared, rvCompared)
		logicBox = append(logicBox, rv.LOperator)
	}

	// 防止溢出
	if len(logicBox) > 1 && len(logicBox) >= len(compared) {
		logicBox = logicBox[:len(logicBox)-1]
	}

	for k, c := range logicBox {
		if len(compared) < k {
			return false, ParseRuleOutOfRangeErr
		}

		flag, err = operator.LogicCombine(c, compared[k], compared[k+1])
		if err != nil {
			return false, errors.WithStack(err)
		}
		compared[k+1] = flag
	}
	return flag, nil
}

// 根据规则，广告微信数据信息，解析决策
func (p *Parse) ParseDecisions(decisions []*Decision, ad *Ad) error {
	var err error
	fmt.Println("正在解析决策")
	if len(decisions) != 0 {
		for _, decision := range decisions {
			//获取
			value, err := ad.AdInfo.GetValue(decision.Value.Slot)
			if err != nil {
				return errors.WithStack(err)
			}
			// 计算
			result, err := decision.Value.Calc(value, decision.Value.ValueType, decision.Inputs, decision.AOperator)
			if err != nil {
				return errors.WithStack(err)
			}
			// 初始化结构体重新赋值
			err = ad.AdInfo.SetValue(result, decision.Value.Slot)
			if err != nil {
				return errors.WithStack(err)
			}
		}
	} else {
		glog.Error(CouldNotFindParseDecisionsErr.Error())
	}

	return errors.WithStack(err)
}

// 根据规则，广告微信数据信息，解析指令
func (p *Parse) ParseOrders(orders []*Order, ad *Ad) error {
	var err error
	fmt.Println("正在解析指令")
	if len(orders) != 0 {
		for _, order := range orders {
			if order.Handle != nil {
				// 执行指令
				err := order.Handle(ad)
				if err != nil {
					return errors.WithStack(err)
				}
			}
		}
	} else {
		glog.Error(CouldNotFindParseOrdersErr.Error())
	}
	return err
}
