package autodmn

// 决策指令
type OrderType string

type Order struct {
	Handle //  自己是实现 Handle，实现业务绑定
}

type Handle func(adSlotData map[string]interface{}) error

func NewOrder() *Order {
	return &Order{}
}
