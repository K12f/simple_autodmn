package autodmn

var AdSlot = make(map[string]interface{})

func init() {
	AdSlot["AdRuleEnableTime"] = AdRuleEnableTime{}
	AdSlot["AdKeepPutTime"] = AdKeepPutTime{}
	AdSlot["AdLeastTime"] = AdLeastTime{}
	AdSlot["AdSpeedRate"] = AdSpeedRate{}
	AdSlot["AdConvCost"] = AdConvCost{}
	AdSlot["AdCurCost"] = AdCurCost{}
	AdSlot["AdExpoSpeed"] = AdExpoSpeed{}
	AdSlot["AdDayCost"] = AdDayCost{}
	AdSlot["ACR"] = ACR{}
	AdSlot["Account"] = Account{}
}

type AdConfig struct {
	SpID string `json:"sp_id"`   // 服务商Id
	GhID string `json:"gh_id"`   // ghid
	AdID string `json:"ad_id"`   // 广告ID
	Name string `json:"ad_name"` // 广告名称
}

type AdInfo struct {
	AdRuleEnableTime AdRuleEnableTime `json:"ad_rule_enable_time"` // 规则有效时间 自定义
	AdKeepPutTime    AdKeepPutTime    `json:"ad_keep_put_time"`    // 继续投放时长  自定义
	AdLeastTime      AdLeastTime      `json:"ad_least_time"`       //广告起投时长
	AdSpeedRate      AdSpeedRate      `json:"ad_speed_rate"`       //广告消耗速度
	AdConvCost       AdConvCost       `json:"ad_conv_cost"`        //广告转化成本
	AdCurCost        AdCurCost        `json:"ad_cur_cost"`         //广告当前出价
	AdExpoSpeed      AdExpoSpeed      `json:"ad_expo_speed"`       //广告曝光速度
	AdDayCost        AdDayCost        `json:"ad_day_cost"`         //广告当日花费
	ACR              ACR              `json:"ad_cr"`               //广告目标转化率
	Account          Account          `json:"ad_account"`          //账户可用余额
}

type Ad struct {
	// 配置信息
	AdConfig AdConfig `json:"ad_config"`
	// 参数信息
	AdInfo *AdInfo `json:"ad_info"`
	// 广告投发规则
	AdRule *AdRule `json:"ad_rule"`
}

func NewAd() *Ad {
	return &Ad{}
}

func (a *Ad) InjectAdInfoToAdSlot() map[string]interface{} {
	AdSlot["AdRuleEnableTime"] = a.AdInfo.AdRuleEnableTime
	AdSlot["AdKeepPutTime"] = a.AdInfo.AdKeepPutTime
	AdSlot["AdLeastTime"] = a.AdInfo.AdLeastTime
	AdSlot["AdSpeedRate"] = a.AdInfo.AdSpeedRate
	AdSlot["AdConvCost"] = a.AdInfo.AdConvCost
	AdSlot["AdCurCost"] = a.AdInfo.AdCurCost
	AdSlot["AdExpoSpeed"] = a.AdInfo.AdExpoSpeed
	AdSlot["AdDayCost"] = a.AdInfo.AdDayCost
	AdSlot["ACR"] = a.AdInfo.ACR
	AdSlot["Account"] = a.AdInfo.Account
	return AdSlot
}
