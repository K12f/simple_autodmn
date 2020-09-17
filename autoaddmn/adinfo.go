package autoaddmn

import (
	"fmt"
	"github.com/pkg/errors"
)

// 广告info数据map
type AdInfoSlot string

// slot 每个值都有，用来判断该值 对应哪个结构体
const (
	AdRuleEnableTimeSlot AdInfoSlot = "AdRuleEnableTime"
	AdKeepPutTimeSlot    AdInfoSlot = "AdKeepPutTime"
	AdLeastTimeSlot      AdInfoSlot = "AdLeastTime"
	AdSpeedRateSlot      AdInfoSlot = "AdSpeedRate"
	AdConvCostSlot       AdInfoSlot = "AdConvCost"
	AdExpoSpeedSlot      AdInfoSlot = "AdExpoSpeed"
	AdDayCostSlot        AdInfoSlot = "AdDayCost"
	ACRSlot              AdInfoSlot = "ACR"
	AccountSlot          AdInfoSlot = "Account"
	AdCurCostSlot        AdInfoSlot = "AdCurCost"
)

// 广告微信平台的数据信息
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

// 规则有效时间
type AdRuleEnableTime struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// 继续投放时长
type AdKeepPutTime struct {
	Value float64 `json:"value"` // 小时
}

// 广告起投时长
type AdLeastTime struct {
	Value float64 `json:"value"` // 小时
}

// 广告消耗速度
type AdSpeedRate struct {
	Value int `json:"value"` //广告消耗速度
}

// 广告转化成本
type AdConvCost struct {
	Value float64 `json:"value"` //价格/分
}

// 广告当前出价
type AdCurCost struct {
	Value float64 `json:"value"` //价格/分
}

// 广告曝光速度
type AdExpoSpeed struct {
	Value int `json:"value"` //
}

// 广告当日花费
type AdDayCost struct {
	Value float64 `json:"value"` //价格/分
}

// 广告目标转化率
type ACR struct {
	Value float64 `json:"value"` // 百分比
}

// 账户可用余额
type Account struct {
	Value float64 `json:"value"` //价格/分
}

func NewAdInfo() *AdInfo {
	return &AdInfo{}
}

// 根据slot的类型，适配对应的 结构体，并根据值重新初始化，并设置值
func (ai *AdInfo) SetValue(value []interface{}, slot AdInfoSlot) error {
	switch slot {
	case AdRuleEnableTimeSlot:
		ai.AdRuleEnableTime = AdRuleEnableTime{
			Start: value[0].(string),
			End:   value[1].(string),
		}
	case AdKeepPutTimeSlot:
		ai.AdKeepPutTime = AdKeepPutTime{
			Value: value[0].(float64),
		}
	case AdLeastTimeSlot:
		ai.AdLeastTime = AdLeastTime{
			Value: value[0].(float64),
		}
	case AdConvCostSlot:
		ai.AdConvCost = AdConvCost{
			Value: value[0].(float64),
		}
	case AdCurCostSlot:
		ai.AdCurCost = AdCurCost{
			Value: value[0].(float64),
		}
	case AdDayCostSlot:
		ai.AdDayCost = AdDayCost{
			Value: value[0].(float64),
		}
	case ACRSlot:
		ai.ACR = ACR{
			Value: value[0].(float64),
		}
	case AccountSlot:
		ai.Account = Account{
			Value: value[0].(float64),
		}
	case AdSpeedRateSlot:
		ai.AdSpeedRate = AdSpeedRate{
			Value: value[0].(int),
		}
	case AdExpoSpeedSlot:
		ai.AdExpoSpeed = AdExpoSpeed{
			Value: value[0].(int),
		}
	default:
		return errors.WithStack(fmt.Errorf("未知的类型"))
	}
	return nil
}

// 根据v的类型，适配结构体，获取到对应的值并返回
func (ai *AdInfo) GetValue(slot AdInfoSlot) ([]interface{}, error) {
	var value []interface{}
	var err error
	switch slot {
	case AdRuleEnableTimeSlot:
		value = append(value, ai.AdRuleEnableTime.Start)
		value = append(value, ai.AdRuleEnableTime.End)
	case AdKeepPutTimeSlot:
		value = append(value, ai.AdKeepPutTime.Value)
	case AdLeastTimeSlot:
		value = append(value, ai.AdLeastTime.Value)
	case AdConvCostSlot:
		value = append(value, ai.AdConvCost.Value)
	case AdCurCostSlot:
		value = append(value, ai.AdCurCost.Value)
	case AdDayCostSlot:
		value = append(value, ai.AdDayCost.Value)
	case ACRSlot:
		value = append(value, ai.ACR.Value)
	case AccountSlot:
		value = append(value, ai.Account.Value)
	case AdSpeedRateSlot:
		value = append(value, ai.AdSpeedRate.Value)
	case AdExpoSpeedSlot:
		value = append(value, ai.AdExpoSpeed.Value)
	default:
		err = errors.WithStack(fmt.Errorf("未知的类型"))
	}
	return value, err
}
