package autodmn

type Collector struct {
	Ad
}

func NewCollector() *Collector {
	return &Collector{}
}

// fake 爬取 一些数据

func (c Collector) Sync(ad AdConfig) *AdInfo {
	strTime := "2020-09-09 02:02:02"

	end := "2020-11-09 02:02:02"

	return &AdInfo{
		AdRuleEnableTime: AdRuleEnableTime{Start: strTime, End: end},
		AdKeepPutTime:    AdKeepPutTime{Value: 100},
		AdLeastTime:      AdLeastTime{Value: 200},
		AdSpeedRate:      AdSpeedRate{Value: 300},
		AdConvCost:       AdConvCost{Value: 400},
		AdCurCost:        AdCurCost{Value: 500},
		AdExpoSpeed:      AdExpoSpeed{Value: 600},
		AdDayCost:        AdDayCost{Value: 700},
		ACR:              ACR{Value: 800},
		Account:          Account{Value: 900},
	}
}
