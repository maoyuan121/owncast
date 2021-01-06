package metrics

import (
	"time"
)

// 多久更新一次指标度量
const metricsPollingInterval = 1 * time.Minute

// CollectedMetrics stores different collected + timestamped values.
type CollectedMetrics struct {
	CPUUtilizations  []timestampedValue `json:"cpu"`
	RAMUtilizations  []timestampedValue `json:"memory"`
	DiskUtilizations []timestampedValue `json:"disk"`

	Viewers []timestampedValue `json:"-"`
}

// Metrics is the shared Metrics instance.
var Metrics *CollectedMetrics

// Start 将开始收集度量指标，并根据情况选择报警
func Start() {
	Metrics = new(CollectedMetrics)
	go startViewerCollectionMetrics()
	handlePolling()

	for range time.Tick(metricsPollingInterval) {
		handlePolling()
	}
}

func handlePolling() {
	// 收集硬件状态
	collectCPUUtilization()
	collectRAMUtilization()
	collectDiskUtilization()

	// 报警
	handleAlerting()
}
