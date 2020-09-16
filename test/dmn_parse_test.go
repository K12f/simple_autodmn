package test

import (
	"fmt"
	"simple_auto_dmn/autodmn"
	"testing"
)

func TestRules(t *testing.T) {

	// 广告配置
	adConfig := autodmn.AdConfig{
		SpID: "12",
		GhID: "12",
		AdID: "123",
		Name: "广告名",
	}

	// adrule
	adrule := autodmn.NewAdRule()

	component := FakeRuleComponent()

	adrule.Set(component)

	ad := autodmn.NewAd()

	ad.AdRule = adrule
	ad.AdConfig = adConfig

	kernel := autodmn.NewKernel()
	err := kernel.Startup(ad)
	if err != nil {
		fmt.Println(err)
	}
}

func FakeRuleComponent() *autodmn.Components {

	//left

	//rule
	leftRule1 := &autodmn.Rule{
		Value: autodmn.Value{
			Name:      "规则有效时间",
			Desc:      "设置当前规则起始时间",
			Slot:      "AdRuleEnableTime",
			ValueType: autodmn.OMD,
		},
		Operator: autodmn.BETWEEN,
		Inputs: []autodmn.Input{{
			Value: "2020-09-08 12:12:12",
		}, {
			Value: "2020-12-12 12:12:12",
		}},
		Condition: autodmn.AND,
	}

	leftRule2 := &autodmn.Rule{
		Value: autodmn.Value{
			Name:      "广告消耗速度",
			Desc:      "广告消耗速度",
			Slot:      "AdSpeedRate",
			ValueType: autodmn.OI,
		},
		Operator: autodmn.GT,
		Inputs: []autodmn.Input{{
			Value: "12",
		}},
	}

	// decision
	leftDecision1 := &autodmn.Decision{
		Value: autodmn.Value{
			Name:      "广告当前出价",
			Desc:      "广告当前出价",
			Slot:      "AdCurCost",
			ValueType: autodmn.OF,
		},
		Operator: autodmn.ADD,
		Inputs: []autodmn.Input{{
			Value: 1000,
		}},
	}
	leftDecision2 := &autodmn.Decision{
		Value: autodmn.Value{
			Name:      "广告当前出价",
			Desc:      "广告当前出价",
			Slot:      "AdCurCost",
			ValueType: autodmn.OF,
		},
		Operator: autodmn.SUB,
		Inputs: []autodmn.Input{{
			Value: 2000,
		}},
	}
	// order
	leftOrder1 := &autodmn.Order{
		Handle: func(adSlotData map[string]interface{}) (err error) {
			fmt.Println(adSlotData)
			fmt.Println("left order")
			return err
		},
	}
	//
	//right

	// rightDecision1
	rightDecision1 := &autodmn.Decision{
		Value: autodmn.Value{
			Name:      "广告当前出价",
			Desc:      "广告当前出价",
			Slot:      "AdCurCost",
			ValueType: autodmn.OF,
		},
		Operator: autodmn.SUB,
		Inputs: []autodmn.Input{{
			Value: 100,
		}},
	}
	rightOrder1 := &autodmn.Order{
		Handle: func(adSlotData map[string]interface{}) (err error) {
			fmt.Println(adSlotData)
			fmt.Println("right order")
			return err
		},
	}
	//bottom

	bottomOrder := &autodmn.Order{
		Handle: func(adSlotData map[string]interface{}) (err error) {
			fmt.Println(adSlotData)
			fmt.Println("bottom order")
			return err
		},
	}

	components := autodmn.NewComponents()
	componentsLeft := autodmn.NewComponentsLeft()
	componentsRight := autodmn.NewComponentsRight()

	componentsLeft.PushRules(leftRule1)
	componentsLeft.PushRules(leftRule2)
	componentsLeft.PushDecisions(leftDecision1)
	componentsLeft.PushDecisions(leftDecision2)
	componentsLeft.PushOrders(leftOrder1)

	//components2 := autodmn.NewComponents()
	//componentsLeft2 := autodmn.NewComponentsLeft()
	//componentsLeft2.PushRules(leftRule1)
	//componentsLeft2.PushRules(leftRule2)
	//componentsLeft2.PushDecisions(leftDecision1)
	//componentsLeft2.PushOrders(leftOrder1)
	//components2.PushLeft(componentsLeft2)
	//
	//componentsLeft.PushComponents(components2)

	//
	componentsRight.PushOrders(rightOrder1)
	componentsRight.PushDecisions(rightDecision1)

	components.PushLeft(componentsLeft)
	components.PushRight(componentsRight)
	components.PushBottom(bottomOrder)

	return components
}
