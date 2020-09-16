package autodmn

// 结构图
//
//
//
//

type AdRule struct {
	Components []*Components `json:"components"` //规则组成
}

func NewAdRule() *AdRule {
	return &AdRule{}
}

func (ar *AdRule) Set(component *Components) {
	ar.Components = append(ar.Components, component)
}

func (ar *AdRule) Get() []*Components {
	return ar.Components
}
