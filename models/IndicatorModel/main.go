package IndicatorModel

import "time"

type Indicator struct {
	ID        int64                  `json:"id"`
	Name      string                 `json:"name"`
	Weight    int8                   `json:"weight"`
	Value     map[string]interface{} `json:"value"`
	Formula   string `json:"formula"`
	CreatedAt time.Time `json:"created_at"`
}