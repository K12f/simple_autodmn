package autoaddmn

// 规则
type Rule struct {
	Value     Value              `json:"value"`            // 值
	COperator ComparisonOperator `json:"compare_operator"` // 比较操作符
	Inputs    []Input            `json:"inputs"`           // 输入
	LOperator LogicOperatorType  `json:"logic_operator"`   // 逻辑操作符
}

func NewRule() *Rule {
	return &Rule{}
}
