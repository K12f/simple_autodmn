package autoaddmn

type Order struct {
	Handle `json:"handle,-"` //  自己是实现 Handle，实现业务绑定
}

// 接收 Ad 广告数据
type Handle func(ad *Ad) error
