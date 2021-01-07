package metrics

import (
	"time"
)

// 多久更新一次指标度量
const metricsPollingInterval = 1 * time.Minute

// CollectedMetrics 存储各种度量指标
type CollectedMetrics struct {
	CPUUtilizations  []timestampedValue `json:"cpu"`    // cpu 度量指标
	RAMUtilizations  []timestampedValue `json:"memory"` // 内存度量指标
	DiskUtilizations []timestampedValue `json:"disk"`   // 硬盘度量指标

	Viewers []timestampedValue `json:"-"` // 在线人数度量指标
}

// Metrics 是一个共享的（引用类型，指针） Metrics 实例
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
