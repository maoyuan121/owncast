package metrics

import "time"

// 度量条目
type timestampedValue struct {
	Time  time.Time `json:"time"`  // 时间点
	Value int       `json:"value"` // 指标值
}
