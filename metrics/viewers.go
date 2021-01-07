package metrics

import (
	"time"

	"github.com/owncast/owncast/core"
)

// 多久统计一次在线人数
const viewerMetricsPollingInterval = 2 * time.Minute

// 用一个新的协程调用，开始定时统计在线人数
func startViewerCollectionMetrics() {
	collectViewerCount()

	for range time.Tick(viewerMetricsPollingInterval) {
		collectViewerCount()
	}
}

// 收集在线人数度量指标
// 如果超过最大收集数，那么删除最老的一个度量指标
func collectViewerCount() {
	if len(Metrics.Viewers) > maxCollectionValues {
		Metrics.Viewers = Metrics.Viewers[1:]
	}

	count := core.GetStatus().ViewerCount
	value := timestampedValue{
		Value: count,
		Time:  time.Now(),
	}
	Metrics.Viewers = append(Metrics.Viewers, value)
}
