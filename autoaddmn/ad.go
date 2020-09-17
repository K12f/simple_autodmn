package autoaddmn

// 一条广告
type Ad struct {
	// 配置信息
	AdConfig AdConfig `json:"ad_config"`
	// 参数信息
	AdInfo *AdInfo `json:"ad_info"`
	// 广告投发规则
	AdRule *AdRule `json:"ad_rule"`
}

// 广告配置信息
type AdConfig struct {
	SpID string `json:"sp_id"`   // 服务商Id
	GhID string `json:"gh_id"`   // ghid
	AdID string `json:"ad_id"`   // 广告ID
	Name string `json:"ad_name"` // 广告名称
}

// new
func NewAd() *Ad {
	return &Ad{}
}
