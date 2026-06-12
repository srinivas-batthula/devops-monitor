package collector

import (
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

func GetOSMetrics() (string, int, int, float64, float64, float64) {
	osType := runtime.GOOS              // Dynamic: "windows", "linux", or "darwin"
	goroutines := runtime.NumGoroutine()
	cpuCores := runtime.NumCPU()

	// 1. Fetch Real CPU Utilization Percentage
	// We sample the CPU cycles over a 100ms interval to calculate usage
	cpuPercentages, err := cpu.Percent(100*time.Millisecond, false)
	var cpuUsage float64
	if err == nil && len(cpuPercentages) > 0 {
		cpuUsage = cpuPercentages[0]
	}

	// 2. Fetch Real Virtual Memory Utilization Percentage
	vMem, err := mem.VirtualMemory()
	var memUsage float64
	if err == nil {
		memUsage = vMem.UsedPercent
	}

	// 3. Fetch Real Disk Storage Utilization Percentage
	// Use "/" for Linux/Mac roots, gopsutil automatically resolves it to "C:" on Windows
	diskStats, err := disk.Usage("/")
	var diskUsage float64
	if err == nil {
		diskUsage = diskStats.UsedPercent
	}

	return osType, goroutines, cpuCores, cpuUsage, memUsage, diskUsage
}