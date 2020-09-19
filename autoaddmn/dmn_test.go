package autoaddmn

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"testing"
)

func TestRules(t *testing.T) {
	//1.实例化ad
	ad := NewAd()

	//1.fake 广告配置
	ad.AdConfig = AdConfig{
		SpID: "12",
		GhID: "12",
		AdID: "123",
		Name: "广告名",
	}

	//2.抓取数据
	ad.AdInfo = &AdInfo{
		AdRuleEnableTime: AdRuleEnableTime{Start: "2020-09-09 02:02:02", End: "2020-11-09 02:02:02"},
		AdKeepPutTime:    AdKeepPutTime{Value: 100},
		AdLeastTime:      AdLeastTime{Value: 200},
		AdSpeedRate:      AdSpeedRate{Value: 300},
		AdConvCost:       AdConvCost{Value: 400},
		AdCurCost:        AdCurCost{Value: 500},
		AdExpoSpeed:      AdExpoSpeed{Value: 600},
		AdDayCost:        AdDayCost{Value: 700},
		ACR:              ACR{Value: 800},
		Account:          Account{Value: 900},
	}

	// 3.获取规则数据
	adrule := NewAdRule()

	// 从数据库中获取，这里是fake一些数据
	// 将规则数据放入组件中
	component := FakeRuleComponent()

	adrule.Set(component)

	ad.AdRule = adrule

	j, _ := gjson.Encode(ad)

	fmt.Println(string(j))
	//4. 开始解析
	kernel := NewKernel()
	err := kernel.Startup(ad)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("------", ad.AdInfo)
	//var c []*Components
	//c = append(c, component)
	//err = kernel.handle(c, ad)
	//if err != nil {
	//	fmt.Println("出现错误了")
	//}
}

func FakeRuleComponent() *Components {
	//left

	////rule
	leftRule1 := &Rule{
		Value: Value{
			//Name:      "规则有效时间",
			//Desc:      "设置当前规则起始时间",
			Slot:      AdRuleEnableTimeSlot,
			ValueType: OMD,
		},
		COperator: BETWEEN,
		Inputs: []Input{{
			Value: "2020-09-08 12:12:12",
		}, {
			Value: "2020-12-12 12:12:12",
		}},
		LOperator: AND,
	}

	leftRule2 := &Rule{
		Value: Value{
			Slot:      AdSpeedRateSlot,
			ValueType: OF,
		},
		COperator: GT,
		Inputs: []Input{{
			Value: 12,
		}},
		LOperator: AND,
	}

	leftRule3 := &Rule{
		Value: Value{
			Slot:      AdLeastTimeSlot,
			ValueType: OF,
		},
		COperator: LET,
		Inputs: []Input{{
			Value: 10,
		}},
		LOperator: AND,
	}

	leftRule4 := &Rule{
		Value: Value{
			Slot:      AdConvCostSlot,
			ValueType: OF,
		},
		COperator: LET,
		Inputs: []Input{{
			Value: 10,
		}},
		LOperator: AND,
	}

	leftRule5 := &Rule{
		Value: Value{
			Slot:      AdCurCostSlot,
			ValueType: OF,
		},
		COperator: LET,
		Inputs: []Input{{
			Value: 10,
		}},
		LOperator: AND,
	}

	leftRule6 := &Rule{
		Value: Value{
			Slot:      AdExpoSpeedSlot,
			ValueType: OF,
		},
		COperator: LET,
		Inputs: []Input{{
			Value: 10,
		}},
		LOperator: AND,
	}

	leftRule7 := &Rule{
		Value: Value{
			Slot:      AdDayCostSlot,
			ValueType: OF,
		},
		COperator: LET,
		Inputs: []Input{{
			Value: 10,
		}},
		LOperator: AND,
	}
	leftRule8 := &Rule{
		Value: Value{
			Slot:      ACRSlot,
			ValueType: OF,
		},
		COperator: LET,
		Inputs: []Input{{
			Value: 10,
		}},
		LOperator: OR,
	}

	leftRule9 := &Rule{
		Value: Value{
			Slot:      AccountSlot,
			ValueType: OF,
		},
		COperator: GT,
		Inputs: []Input{{
			Value: 10,
		}},
		LOperator: AND,
	}

	// decision
	leftDecision1 := &Decision{
		Value: Value{
			Slot:      AdCurCostSlot,
			ValueType: OF,
		},
		AOperator: ADD,
		Inputs: []Input{{
			Value: 1000,
		}},
	}

	leftDecision2 := &Decision{
		Value: Value{
			Slot:      AdKeepPutTimeSlot,
			ValueType: OF,
		},
		AOperator: Per,
		Inputs: []Input{{
			Value: .5,
		}},
	}

	leftDecision3 := &Decision{
		Value: Value{
			Slot:      AdRuleEnableTimeSlot,
			ValueType: OMD,
		},
		AOperator: SUB,
		Inputs: []Input{{
			Value: 20,
		}},
	}

	leftDecision4 := &Decision{
		Value: Value{
			Slot:      AdLeastTimeSlot,
			ValueType: OF,
		},
		AOperator: SUB,
		Inputs: []Input{{
			Value: 2000,
		}},
	}

	leftDecision5 := &Decision{
		Value: Value{
			Slot:      AdSpeedRateSlot,
			ValueType: OF,
		},
		AOperator: SUB,
		Inputs: []Input{{
			Value: 2000,
		}},
	}

	leftDecision6 := &Decision{
		Value: Value{
			Slot:      AdConvCostSlot,
			ValueType: OF,
		},
		AOperator: SUB,
		Inputs: []Input{{
			Value: 2000,
		}},
	}

	leftDecision7 := &Decision{
		Value: Value{
			Slot:      AdCurCostSlot,
			ValueType: OF,
		},
		AOperator: SUB,
		Inputs: []Input{{
			Value: 2000,
		}},
	}

	leftDecision8 := &Decision{
		Value: Value{
			Slot:      AdExpoSpeedSlot,
			ValueType: OF,
		},
		AOperator: SUB,
		Inputs: []Input{{
			Value: 2000,
		}},
	}

	leftDecision9 := &Decision{
		Value: Value{
			Slot:      AdDayCostSlot,
			ValueType: OF,
		},
		AOperator: ADD,
		Inputs: []Input{{
			Value: 2000,
		}},
	}

	leftDecision10 := &Decision{
		Value: Value{
			Slot:      ACRSlot,
			ValueType: OF,
		},
		AOperator: Div,
		Inputs: []Input{{
			Value: 23,
		}},
	}

	leftDecision11 := &Decision{
		Value: Value{
			Slot:      AccountSlot,
			ValueType: OF,
		},
		AOperator: Mul,
		Inputs: []Input{{
			Value: 12,
		}},
	}
	// order
	leftOrder1 := &Order{
		Handle: func(ad *Ad) (err error) {
			fmt.Println(ad)
			fmt.Println(ad.AdInfo)
			fmt.Println("left order")
			return err
		},
	}
	//
	//right

	// rightDecision1
	rightDecision1 := &Decision{
		Value: Value{
			Slot:      "AdCurCost",
			ValueType: OF,
		},
		AOperator: SUB,
		Inputs: []Input{{
			Value: 100,
		}},
	}
	rightOrder1 := &Order{
		Handle: func(ad *Ad) (err error) {
			fmt.Println(ad)
			fmt.Println("right order")
			return err
		},
	}
	//bottom

	bottomOrder := &Order{
		Handle: func(ad *Ad) (err error) {
			fmt.Println(ad)
			fmt.Println("bottom order")
			return err
		},
	}

	components := NewComponents()
	componentsLeft := NewComponentsLeft()
	componentsRight := NewComponentsRight()
	componentsBottom := NewComponentsBottom()

	componentsLeft.PushRules(leftRule1, leftRule2, leftRule3, leftRule4, leftRule5, leftRule6, leftRule7, leftRule8, leftRule9)

	componentsLeft.PushDecisions(leftDecision1, leftDecision2, leftDecision3, leftDecision4,
		leftDecision5, leftDecision6, leftDecision7, leftDecision8, leftDecision9, leftDecision10, leftDecision11)

	componentsLeft.PushOrders(leftOrder1)

	components2 := NewComponents()
	componentsLeft2 := NewComponentsLeft()
	componentsLeft2.PushRules(leftRule1, leftRule2)
	componentsLeft2.PushDecisions(leftDecision1)
	componentsLeft2.PushOrders(leftOrder1)
	components2.PushLeft(componentsLeft2)

	componentsLeft.PushComponents(components2)

	//
	componentsRight.PushOrders(rightOrder1)
	componentsRight.PushDecisions(rightDecision1)

	componentsBottom.PushOrders(bottomOrder)

	components.PushLeft(componentsLeft)
	components.PushRight(componentsRight)
	components.PushBottom(componentsBottom)

	return components
}
