package autoaddmn

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func TestAdRuleJson(t *testing.T) {
	data := `{
  "ad_config": {
    "sp_id": "12",
    "gh_id": "12",
    "ad_id": "123",
    "ad_name": "广告名"
  },
  "ad_info": {
    "ad_rule_enable_time": {
      "start": "2020-09-09 02:02:02",
      "end": "2020-11-09 02:02:02"
    },
    "ad_keep_put_time": {
      "value": 100
    },
    "ad_least_time": {
      "value": 200
    },
    "ad_speed_rate": {
      "value": 300
    },
    "ad_conv_cost": {
      "value": 400
    },
    "ad_cur_cost": {
      "value": 500
    },
    "ad_expo_speed": {
      "value": 600
    },
    "ad_day_cost": {
      "value": 700
    },
    "ad_cr": {
      "value": 800
    },
    "ad_account": {
      "value": 900
    }
  },
  "ad_rule": {
    "components": [
      {
        "left": {
          "rules": [
            {
              "value": {
                "slot": "AdRuleEnableTime",
                "value": null,
                "value_type": "mul_date"
              },
              "compare_operator": "in",
              "inputs": [
                {
                  "value": "2020-09-08 12:12:12"
                },
                {
                  "value": "2020-12-12 12:12:12"
                }
              ],
              "logic_operator": "&&"
            },
            {
              "value": {
                "slot": "AdSpeedRate",
                "value": null,
                "value_type": "float64"
              },
              "compare_operator": ">",
              "inputs": [
                {
                  "value": "12"
                }
              ],
              "logic_operator": "&&"
            },
            {
              "value": {
                "slot": "AdLeastTime",
                "value": null,
                "value_type": "float64"
              },
              "compare_operator": "<=",
              "inputs": [
                {
                  "value": 10
                }
              ],
              "logic_operator": "&&"
            },
            {
              "value": {
                "slot": "AdConvCost",
                "value": null,
                "value_type": "float64"
              },
              "compare_operator": "<=",
              "inputs": [
                {
                  "value": 10
                }
              ],
              "logic_operator": "&&"
            },
            {
              "value": {
                "slot": "AdCurCost",
                "value": null,
                "value_type": "float64"
              },
              "compare_operator": "<=",
              "inputs": [
                {
                  "value": 10
                }
              ],
              "logic_operator": "&&"
            },
            {
              "value": {
                "slot": "AdExpoSpeed",
                "value": null,
                "value_type": "float64"
              },
              "compare_operator": "<=",
              "inputs": [
                {
                  "value": 10
                }
              ],
              "logic_operator": "&&"
            },
            {
              "value": {
                "slot": "AdDayCost",
                "value": null,
                "value_type": "float64"
              },
              "compare_operator": "<=",
              "inputs": [
                {
                  "value": 10
                }
              ],
              "logic_operator": "&&"
            },
            {
              "value": {
                "slot": "ACR",
                "value": null,
                "value_type": "float64"
              },
              "compare_operator": "<=",
              "inputs": [
                {
                  "value": 10
                }
              ],
              "logic_operator": "||"
            },
            {
              "value": {
                "slot": "Account",
                "value": null,
                "value_type": "float64"
              },
              "compare_operator": ">",
              "inputs": [
                {
                  "value": 10
                }
              ],
              "logic_operator": "&&"
            }
          ],
          "decisions": [
            {
              "value": {
                "slot": "AdCurCost",
                "value": null,
                "value_type": "float64"
              },
              "ari_operator": "+",
              "inputs": [
                {
                  "value": 1000
                }
              ]
            },
            {
              "value": {
                "slot": "AdKeepPutTime",
                "value": null,
                "value_type": "float64"
              },
              "ari_operator": "%",
              "inputs": [
                {
                  "value": 0.5
                }
              ]
            },
            {
              "value": {
                "slot": "AdRuleEnableTime",
                "value": null,
                "value_type": "mul_date"
              },
              "ari_operator": "-",
              "inputs": [
                {
                  "value": 20
                },
                {
                  "value": 50
                }
              ]
            },
            {
              "value": {
                "slot": "AdLeastTime",
                "value": null,
                "value_type": "float64"
              },
              "ari_operator": "-",
              "inputs": [
                {
                  "value": 2000
                }
              ]
            },
            {
              "value": {
                "slot": "AdSpeedRate",
                "value": null,
                "value_type": "float64"
              },
              "ari_operator": "-",
              "inputs": [
                {
                  "value": 2000
                }
              ]
            },
            {
              "value": {
                "slot": "AdConvCost",
                "value": null,
                "value_type": "float64"
              },
              "ari_operator": "-",
              "inputs": [
                {
                  "value": 2000
                }
              ]
            },
            {
              "value": {
                "slot": "AdCurCost",
                "value": null,
                "value_type": "float64"
              },
              "ari_operator": "-",
              "inputs": [
                {
                  "value": 2000
                }
              ]
            },
            {
              "value": {
                "slot": "AdExpoSpeed",
                "value": null,
                "value_type": "float64"
              },
              "ari_operator": "-",
              "inputs": [
                {
                  "value": 2000
                }
              ]
            },
            {
              "value": {
                "slot": "AdDayCost",
                "value": null,
                "value_type": "float64"
              },
              "ari_operator": "+",
              "inputs": [
                {
                  "value": 2000
                }
              ]
            },
            {
              "value": {
                "slot": "ACR",
                "value": null,
                "value_type": "float64"
              },
              "ari_operator": "/",
              "inputs": [
                {
                  "value": 23
                }
              ]
            },
            {
              "value": {
                "slot": "Account",
                "value": null,
                "value_type": "float64"
              },
              "ari_operator": "*",
              "inputs": [
                {
                  "value": 12
                }
              ]
            }
          ],
          "orders": [
            {}
          ],
          "components": [
            {
              "left": {
                "rules": [
                  {
                    "value": {
                      "slot": "AdRuleEnableTime",
                      "value": null,
                      "value_type": "mul_date"
                    },
                    "compare_operator": "in",
                    "inputs": [
                      {
                        "value": "2020-09-08 12:12:12"
                      },
                      {
                        "value": "2020-12-12 12:12:12"
                      }
                    ],
                    "logic_operator": "&&"
                  },
                  {
                    "value": {
                      "slot": "AdSpeedRate",
                      "value": null,
                      "value_type": "float64"
                    },
                    "compare_operator": ">",
                    "inputs": [
                      {
                        "value": "12"
                      }
                    ],
                    "logic_operator": "&&"
                  }
                ],
                "decisions": [
                  {
                    "value": {
                      "slot": "AdCurCost",
                      "value": null,
                      "value_type": "float64"
                    },
                    "ari_operator": "+",
                    "inputs": [
                      {
                        "value": 1000
                      }
                    ]
                  }
                ],
                "orders": [
                  {}
                ],
                "components": null
              },
              "right": null,
              "bottom": null
            }
          ]
        },
        "right": {
          "decisions": [
            {
              "value": {
                "slot": "AdCurCost",
                "value": null,
                "value_type": "float64"
              },
              "ari_operator": "-",
              "inputs": [
                {
                  "value": 100
                }
              ]
            }
          ],
          "orders": [
            {}
          ],
          "components": null
        },
        "bottom": {
          "orders": [
            {}
          ]
        }
      }
    ]
  }
}`
	var ad = NewAd()
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(data), ad)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ad.AdInfo)
	kernel := NewKernel()
	err = kernel.Startup(ad)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ad.AdInfo)
}
