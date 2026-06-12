package collector

import (
	"os/exec"
	"strings"
)

func GetDockerMetrics() (int, int) {
	// Execute 'docker ps -q' to count running containers
	cmdRunning := exec.Command("docker", "ps", "-q")
	outRun, err := cmdRunning.Output()
	if err != nil {
		return 0, 0 // Return zeroed out if Docker isn't running/installed
	}
	runningCount := len(strings.Split(strings.TrimSpace(string(outRun)), "\n"))
	if strings.TrimSpace(string(outRun)) == "" { runningCount = 0 }

	// Execute 'docker ps -a -q -f status=exited' to isolate stopped container counts
	cmdStopped := exec.Command("docker", "ps", "-a", "-q", "-f", "status=exited")
	outStop, _ := cmdStopped.Output()
	stoppedCount := len(strings.Split(strings.TrimSpace(string(outStop)), "\n"))
	if strings.TrimSpace(string(outStop)) == "" { stoppedCount = 0 }

	return runningCount, stoppedCount
}