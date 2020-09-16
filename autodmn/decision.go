package autodmn

// 决策
type Decision struct {
	Value    Value              `json:"value"`        // 值
	Operator ArithmeticOperator `json:"ari_operator"` //操作符
	Inputs   []Input            `json:"inputs"`       // 输入
}

func NewDecision() *Decision {
	return &Decision{}
}

func ParseDecision(a, b interface{}, co ComparisonOperator) {
	switch co {

	}
}

//func GetType(value interface{})  {
//    switch value.(type) {
//    case int:
//        return "int"
//    case float64:
//    case string:
//    case time.Time:
//    default:
//        panic("error")
//    }
//}

func (d *Decision) Fake1() *Decision {
	return &Decision{
		Value: Value{
			Name: "广告当前出价",
			Desc: "广告当前出价",
			Slot: "AdPrice",
			//Value: value,
		},
		Operator: ADD,
		Inputs: []Input{{
			Value: 1000.00,
		}},
	}
}
