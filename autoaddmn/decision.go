package autoaddmn

// 决策
type Decision struct {
	Value     Value              `json:"value"`        // 值
	AOperator ArithmeticOperator `json:"ari_operator"` // 算数运算符
	Inputs    []Input            `json:"inputs"`       // 输入
}
