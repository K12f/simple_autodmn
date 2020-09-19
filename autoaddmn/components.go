package autoaddmn

/**
组件
*/
type Components struct {
	Left   *Left   `json:"left"`   // 左节点，即 if
	Right  *Right  `json:"right"`  // 右节点， 即else
	Bottom *Bottom `json:"bottom"` // 底部节点， 即 最后指令
}

// 左结构体, if
type Left struct {
	Rules      []*Rule       `json:"rules"`      // 规则
	Decisions  []*Decision   `json:"decisions"`  // 决策
	Orders     []*Order      `json:"orders"`     // 指令
	Components []*Components `json:"components"` // 子组件
}

// 右结构体, else
type Right struct {
	Decisions  []*Decision   `json:"decisions"`  // 决策
	Orders     []*Order      `json:"orders"`     // 指令
	Components []*Components `json:"components"` // 子组件
}

// 右结构体, else
type Bottom struct {
	Orders []*Order `json:"orders"` // 指令
}

func NewComponents() *Components {
	return &Components{}
}

// push数据到 左结构体中
func (c *Components) PushLeft(left *Left) {
	c.Left = left
}

// push数据到 右结构体中
func (c *Components) PushRight(right *Right) {
	c.Right = right
}

// push数据到 右结构体中
func (c *Components) PushBottom(bottom *Bottom) {
	c.Bottom = bottom
}

// 左结构体实例化
func NewLeft() *Left {
	return &Left{}
}

// push 规则数据
func (l *Left) PushRules(rule ...*Rule) {
	l.Rules = append(l.Rules, rule...)
}

// push 决策数据
func (l *Left) PushDecisions(decision ...*Decision) {
	l.Decisions = append(l.Decisions, decision...)
}

// push 指令数据
func (l *Left) PushOrders(order *Order) {
	l.Orders = append(l.Orders, order)
}

// push 左子组件
func (l *Left) PushComponents(component ...*Components) {
	l.Components = append(l.Components, component...)
}

// 右结构体实例化
func NewRight() *Right {
	return &Right{}
}

// push 右结构体决策
func (r *Right) PushDecisions(decision ...*Decision) {
	r.Decisions = append(r.Decisions, decision...)
}

// push 右结构体指令
func (r *Right) PushOrders(order ...*Order) {
	r.Orders = append(r.Orders, order...)
}

// push 右结构体子组件
func (r *Right) PushComponents(component ...*Components) {
	r.Components = append(r.Components, component...)
}

// 右结构体实例化
func NewBottom() *Bottom {
	return &Bottom{}
}

// push 右结构体指令
func (b *Bottom) PushOrders(order ...*Order) {
	b.Orders = append(b.Orders, order...)
}
