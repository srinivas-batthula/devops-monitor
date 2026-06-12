package engine

import (
	"sync"
	"devops-monitor/internal/collector"
	"devops-monitor/internal/models"
)

// RunConcurrentScan orchestrates " parallel collectors " to assemble a system report
func RunConcurrentScan() models.SystemReport {
	var wg sync.WaitGroup
	var mu sync.Mutex
	report := models.SystemReport{}

	// Task 1: Collect OS data via Goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		osType, gCount, cores, cpu, mem, disk := collector.GetOSMetrics()
		
		mu.Lock() // Locking section to safely write into the shared struct
		report.OSType = osType
		report.Goroutines = gCount
		report.CPUCores = cores
		report.CPUUsage = cpu
		report.MemoryUsage = mem
		report.DiskUsage = disk
		mu.Unlock()
	}()

	// Task 2: Collect Docker data via Goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		run, stop := collector.GetDockerMetrics()
		
		mu.Lock()
		report.DockerRunning = run
		report.DockerStopped = stop
		mu.Unlock()
	}()

	// Task 3: Collect Kubernetes data via Goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		nodes, pods := collector.GetK8sMetrics()
		
		mu.Lock()
		report.K8sNodes = nodes
		report.K8sPods = pods
		mu.Unlock()
	}()

	wg.Wait() // Blocks until all 3-concurrent functions call 'defer wg.Done()'

	// Run Health Analyzer Brain Engine
	report.Status = analyzeHealth(report.CPUUsage, report.MemoryUsage)

	return report
}

func analyzeHealth(cpu, mem float64) string {
	if cpu > 85.0 || mem > 90.0 {
		return "CRITICAL"
	} else if cpu > 70.0 || mem > 75.0 {
		return "WARNING"
	}
	return "HEALTHY"
}