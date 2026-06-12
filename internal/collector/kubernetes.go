package collector

import (
	"os/exec"
	"strings"
)

func GetK8sMetrics() (int, int) {
	// Querying Node resource allocations
	cmdNodes := exec.Command("kubectl", "get", "nodes", "-o", "name")
	outNodes, err := cmdNodes.Output()
	if err != nil {
		return 0, 0 // Fallback gracefully if cluster context isn't running
	}
	nodeCount := len(strings.Split(strings.TrimSpace(string(outNodes)), "\n"))

	// Querying total Cluster Pod allocations
	cmdPods := exec.Command("kubectl", "get", "pods", "-A", "-o", "name")
	outPods, _ := cmdPods.Output()
	podCount := len(strings.Split(strings.TrimSpace(string(outPods)), "\n"))

	return nodeCount, podCount
}