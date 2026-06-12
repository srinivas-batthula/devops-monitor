package report

import (
	"devops-monitor/internal/models"
	"fmt"
)

// RenderTerminal Output maps metric objects into an intuitive layout
func RenderTerminal(r models.SystemReport) {
	fmt.Println("\n========================================")
	fmt.Println("         DEV-OPS MONITOR REPORT         ")
	fmt.Println("========================================")
	fmt.Printf("OS Architecture   : %s\n", r.OSType)
	fmt.Printf("Active Goroutines : %d\n", r.Goroutines)
	fmt.Printf("Logical CPU Cores : %d\n", r.CPUCores)
	fmt.Printf("CPU Utilization   : %.2f%%\n", r.CPUUsage)
	fmt.Printf("Memory Utilization: %.2f%%\n", r.MemoryUsage)
	fmt.Printf("Disk Allocation   : %.2f%%\n", r.DiskUsage)
	fmt.Println("----------------------------------------")
	fmt.Println("Docker Container State:")
	fmt.Printf("  Running         : %d\n", r.DockerRunning)
	fmt.Printf("  Stopped         : %d\n", r.DockerStopped)
	fmt.Println("----------------------------------------")
	fmt.Println("Kubernetes Cluster State:")
	fmt.Printf("  Active Nodes    : %d\n", r.K8sNodes)
	fmt.Printf("  Monitored Pods  : %d\n", r.K8sPods)
	fmt.Println("----------------------------------------")
	fmt.Printf("OVERALL ENGINE HEALTH STATUS: %s\n", r.Status)
	fmt.Println("========================================")
}
