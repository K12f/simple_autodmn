package autoaddmn

import (
	"fmt"
	"github.com/gogf/gf/os/glog"
	"github.com/pkg/errors"
	"time"
)

var (
	Name    = "广告规则自动决策OS"
	Version = "0.0.1"
)

func init() {
	fmt.Printf("%s \n 版本:%s \n 时间:%s", Name, Version, time.Now().Format(TimeFormat))
	fmt.Println("----------------------------")
}

// dmn 核心，对外暴露
type Kernel struct {
}

func NewKernel() *Kernel {
	return &Kernel{}
}

// 根据ad 开启检测及执行决策和指令
func (k Kernel) Startup(ad *Ad) error {
	fmt.Println("决策解析开始")
	fmt.Println("----------------------------")
	glog.Info("广告配置", ad.AdConfig)
	glog.Info("广告数据", ad.AdInfo)

	glog.Info("广告规则组件", &ad.AdRule.Components)

	err := k.handle(ad.AdRule.Components, ad)

	if err != nil {
		return err
	}
	fmt.Println("决策解析结束")
	return nil
}

// 处理
func (k Kernel) handle(components []*Components, ad *Ad) error {
	parse := NewParse()
	// 没有组件
	if len(components) <= 0 {
		return NotFoundRuleComponentsErr
	}
	fmt.Println("----------------------------")
	// 遍历
	fmt.Printf("进入组件: \n")
	for _, v := range components {
		// 组件左节点数据 必须有
		if v.Left == nil {
			return NotFoundRuleComponentsLeftErr
		}

		//左节点解析规则
		leftCompared, err := parse.ParseRules(v.Left.Rules, ad)
		if err != nil {
			return errors.Wrap(err, CouldNotParseRulesErr.Error())
		}

		// 左节点规则判断
		if leftCompared {
			fmt.Println("正在解析左节点数据")
			//左节点决策解析
			err = parse.ParseDecisions(v.Left.Decisions, ad)
			if err != nil {
				return errors.Wrap(err, CouldNotParseDecisionsErr.Error())
			}
			// 左节点 指令解析
			err = parse.ParseOrders(v.Left.Orders, ad)
			if err != nil {
				return errors.Wrap(err, CouldNotParseOrdersErr.Error())
			}
			//如果左节点 也有 components
			if len(v.Left.Components) > 0 {
				fmt.Println("----------------------------")
				fmt.Println("正在进入左节点子组件:")
				//迭代
				err = k.handle(v.Left.Components, ad)
				if err != nil {
					return errors.Wrap(err, CouldNotParseSubComponentsErr.Error())
				}
			}
		} else {
			// 判断是否有组件右节点数据
			fmt.Println("正在解析右节点数据")
			if v.Right != nil {
				// 右节点决策解析
				err := parse.ParseDecisions(v.Right.Decisions, ad)
				if err != nil {
					return errors.Wrap(err, CouldNotParseDecisionsErr.Error())
				}
				// 指令解析
				err = parse.ParseOrders(v.Right.Orders, ad)
				if err != nil {
					return errors.Wrap(err, CouldNotParseOrdersErr.Error())
				}
				//如果右节点 也有 components
				if len(v.Right.Components) > 0 {
					fmt.Println("----------------------------")
					fmt.Println("正在进入右节点子组件")
					err = k.handle(v.Right.Components, ad)
					if err != nil {
						return errors.Wrap(err, CouldNotParseSubComponentsErr.Error())
					}
				}
			} else {
				fmt.Printf("!! ID:%s-%s \n", ad.AdConfig.AdID, NotFoundRuleComponentsRightErr.Error())
				//glog.Error(NotFoundRuleComponentsRightErr.Error())
			}
		}
		//底部指令解析
		if v.Bottom != nil {
			err = parse.ParseOrders(v.Bottom.Orders, ad)
			if err != nil {
				return err
			}
		}
	}
	fmt.Println("----------------------------")
	return nil
}
