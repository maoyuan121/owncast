package metrics

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"

	log "github.com/sirupsen/logrus"
)

// 最大收集多少个度量指标
const maxCollectionValues = 500

// 收集 cpu 度量指标
// 如果超过最大收集数，那么删除最老的一个度量指标
func collectCPUUtilization() {
	if len(Metrics.CPUUtilizations) > maxCollectionValues {
		Metrics.CPUUtilizations = Metrics.CPUUtilizations[1:]
	}

	v, err := cpu.Percent(0, false)
	if err != nil {
		log.Errorln(err)
		return
	}

	metricValue := timestampedValue{time.Now(), int(v[0])}
	Metrics.CPUUtilizations = append(Metrics.CPUUtilizations, metricValue)
}

// 收集内存度量指标
// 如果超过最大收集数，那么删除最老的一个度量指标
func collectRAMUtilization() {
	if len(Metrics.RAMUtilizations) > maxCollectionValues {
		Metrics.RAMUtilizations = Metrics.RAMUtilizations[1:]
	}

	memoryUsage, _ := mem.VirtualMemory()
	metricValue := timestampedValue{time.Now(), int(memoryUsage.UsedPercent)}
	Metrics.RAMUtilizations = append(Metrics.RAMUtilizations, metricValue)
}

// 收集硬盘度量指标
// 如果超过最大收集数，那么删除最老的一个度量指标
func collectDiskUtilization() {
	path := "./"
	diskUse, _ := disk.Usage(path)

	if len(Metrics.DiskUtilizations) > maxCollectionValues {
		Metrics.DiskUtilizations = Metrics.DiskUtilizations[1:]
	}

	metricValue := timestampedValue{time.Now(), int(diskUse.UsedPercent)}
	Metrics.DiskUtilizations = append(Metrics.DiskUtilizations, metricValue)
}
