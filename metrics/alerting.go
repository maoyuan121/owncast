package metrics

import (
	log "github.com/sirupsen/logrus"
)

const maxCPUAlertingThresholdPCT = 80  // CPU 占用率达到多少百分比报警
const maxRAMAlertingThresholdPCT = 80  // 内存占用率达到多少百分比报警
const maxDiskAlertingThresholdPCT = 90 // 硬盘占用率达到多少百分比报警

const alertingError = "The %s utilization of %d%% can cause issues with video generation and delivery. Please visit the documentation at http://owncast.online/docs/troubleshooting/ to help troubleshoot this issue."

// handler 报警
func handleAlerting() {
	handleCPUAlerting()
	handleRAMAlerting()
	handleDiskAlerting()
}

// cpu 报警
// 如果只有两个时间点的 cpu 指标的话，直接不处理
func handleCPUAlerting() {
	if len(Metrics.CPUUtilizations) < 2 {
		return
	}

	avg := recentAverage(Metrics.CPUUtilizations)
	if avg > maxCPUAlertingThresholdPCT {
		log.Errorf(alertingError, "CPU", maxCPUAlertingThresholdPCT)
	}
}

// 内存报警
// 如果只有两个时间点的内存指标的话，直接不处理
func handleRAMAlerting() {
	if len(Metrics.RAMUtilizations) < 2 {
		return
	}

	avg := recentAverage(Metrics.RAMUtilizations)
	if avg > maxRAMAlertingThresholdPCT {
		log.Errorf(alertingError, "memory", maxRAMAlertingThresholdPCT)
	}
}

// 硬盘报警
// 如果只有两个时间点的硬盘指标的话，直接不处理
func handleDiskAlerting() {
	if len(Metrics.DiskUtilizations) < 2 {
		return
	}

	avg := recentAverage(Metrics.DiskUtilizations)

	if avg > maxDiskAlertingThresholdPCT {
		log.Errorf(alertingError, "disk", maxRAMAlertingThresholdPCT)
	}
}

// 求时间点集合的值的评价数
// 具体算法：拿最近两个时间点的值求和取平均值
func recentAverage(values []timestampedValue) int {
	return (values[len(values)-1].Value + values[len(values)-2].Value) / 2
}
