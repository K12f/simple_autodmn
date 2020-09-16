package autodmn

// 规则
type Rule struct {
	Value     Value              `json:"value"`     // 值
	Operator  ComparisonOperator `json:"operator"`  // 比较操作符
	Inputs    []Input            `json:"inputs"`    //输入
	Condition LogicOperatorType  `json:"condition"` //条件
}

func NewRule() *Rule {
	return &Rule{}
}
