package autoaddmn

// 广告规则
type AdRule struct {
	Components []*Components `json:"components"` //规则组成
}

// new
func NewAdRule() *AdRule {
	return &AdRule{}
}

// 将规则组件推入
func (ar *AdRule) Set(component *Components) {
	ar.Components = append(ar.Components, component)
}

// 获取规则组件
func (ar *AdRule) Get() []*Components {
	return ar.Components
}
