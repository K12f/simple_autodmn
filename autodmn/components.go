package autodmn

/**
组件
*/
type Components struct {
	Left   *Left    `json:"left"`   // 左节点，即 if
	Right  *Right   `json:"right"`  // 右节点， 即else
	Bottom []*Order `json:"bottom"` // 顶部节点， 即 最后指令
}

type Left struct {
	Rules      []*Rule       `json:"rules"`      // 规则
	Decisions  []*Decision   `json:"decisions"`  // 决策
	Orders     []*Order      `json:"orders"`     // 指令
	Components []*Components `json:"components"` // 子组件
}

type Right struct {
	Decisions  []*Decision   `json:"decisions"`
	Orders     []*Order      `json:"orders"`
	Components []*Components `json:"components"`
}

func NewComponents() *Components {
	return &Components{}
}

func (c *Components) PushLeft(left *Left) {
	c.Left = left
}

func (c *Components) PushRight(right *Right) {
	c.Right = right
}

func (c *Components) PushBottom(order *Order) {
	c.Bottom = append(c.Bottom, order)
}

func NewComponentsLeft() *Left {
	return &Left{}
}

func (l *Left) PushRules(rule *Rule) {
	l.Rules = append(l.Rules, rule)
}

func (l *Left) PushDecisions(decision *Decision) {
	l.Decisions = append(l.Decisions, decision)
}

func (l *Left) PushOrders(order *Order) {
	l.Orders = append(l.Orders, order)
}

func (l *Left) PushComponents(component *Components) {
	l.Components = append(l.Components, component)
}

func NewComponentsRight() *Right {
	return &Right{}
}

func (r *Right) PushDecisions(decision *Decision) {
	r.Decisions = append(r.Decisions, decision)
}

func (r *Right) PushOrders(order *Order) {
	r.Orders = append(r.Orders, order)
}

func (r *Right) PushComponents(component *Components) {
	r.Components = append(r.Components, component)
}
