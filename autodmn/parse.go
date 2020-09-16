package autodmn

import (
	"github.com/pkg/errors"
)

type Parse struct {
}

func NewParse() *Parse {
	return &Parse{}
}

func (p *Parse) ParseRules(rules []*Rule, adSlotData map[string]interface{}) (bool, error) {
	var compared []bool
	var condBox []LogicOperatorType
	var err error
	var flag bool
	operator := NewOperator()

	for _, rv := range rules {
		rv.Value.Value = adSlotData[rv.Value.Slot]
		rvCompared, err := rv.Value.Compare(rv.Inputs, rv.Operator)

		if err != nil {
			return false, errors.WithStack(err)
		}
		compared = append(compared, rvCompared)
		condBox = append(condBox, rv.Condition)
	}
	condBox = p.DropEmptyCond(condBox)

	for k, c := range condBox {
		if len(compared) < k {
			return false, ParseRuleOutOfRangeErr
		}

		flag, err = operator.LogicCombine(c, compared[k], compared[k+1])
		if err != nil {
			return false, err
		}
		compared[k+1] = flag
	}
	return flag, nil
}

func (p *Parse) ParseDecisions(decisions []*Decision, adSlotData map[string]interface{}) error {
	var err error
	for _, decision := range decisions {
		decision.Value.Value = adSlotData[decision.Value.Slot]
		result, err := decision.Value.Calc(decision.Inputs, decision.Operator)
		if err != nil {
			return errors.WithStack(err)
		}
		value, err := decision.Value.SetValue(result)
		if err != nil {
			return errors.WithStack(err)
		}
		adSlotData[decision.Value.Slot] = value
	}
	return errors.WithStack(err)
}

func (p *Parse) ParseOrders(orders []*Order, adSlotData map[string]interface{}) error {
	var err error
	for _, order := range orders {
		if order.Handle != nil {
			err := order.Handle(adSlotData)
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return err
}

func (p *Parse) DropEmptyCond(types []LogicOperatorType) []LogicOperatorType {
	var newConditionTypes []LogicOperatorType
	for _, tp := range types {
		if tp != "" && len(tp) > 0 {
			newConditionTypes = append(newConditionTypes, tp)
		}
	}
	return newConditionTypes
}
