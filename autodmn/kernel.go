package autodmn

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

type Kernel struct {
}

func NewKernel() *Kernel {
	return &Kernel{}
}

func (k *Kernel) Startup(ad *Ad) error {
	//1.获取 ad的配置信息

	//2.根据ad 配置信息抓取 广告 参数

	collector := NewCollector()

	ad.AdInfo = collector.Sync(ad.AdConfig)

	// 将值注入到map中,方便读取
	adSlotData := ad.InjectAdInfoToAdSlot()
	err := k.Handle(ad.AdRule.Components, adSlotData)

	if err != nil {
		glog.Error(fmt.Sprintf("%+v", err))
	}
	//4.根据rule的结果 做出决策和指令

	return nil
}

func (k *Kernel) Handle(components []*Components, adSlotData map[string]interface{}) error {
	parse := NewParse()

	if len(components) <= 0 {
		return NotFoundRuleComponentsErr
	}

	for _, v := range components {
		left := v.Left
		right := v.Right
		bottom := v.Bottom

		// 计算
		leftCompared, err := parse.ParseRules(left.Rules, adSlotData)
		if err != nil {
			return errors.Wrap(err, "左侧规则解析失败")
		}
		if leftCompared {
			//left 决策
			err = parse.ParseDecisions(left.Decisions, adSlotData)
			if err != nil {
				return errors.WithStack(err)
			}
			err = parse.ParseOrders(left.Orders, adSlotData)
			if err != nil {
				return errors.WithStack(err)
			}
			//如果左侧 也有 components
			if len(left.Components) > 0 {
				return k.Handle(left.Components, adSlotData)
			}
		} else {
			err := parse.ParseDecisions(right.Decisions, adSlotData)
			if err != nil {
				return errors.WithStack(err)
			}
			err = parse.ParseOrders(right.Orders, adSlotData)
			if err != nil {
				return errors.WithStack(err)
			}
			//如果右侧 也有 components
			if len(right.Components) > 0 {
				return k.Handle(right.Components, adSlotData)
			}
		}
		err = parse.ParseOrders(bottom, adSlotData)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
