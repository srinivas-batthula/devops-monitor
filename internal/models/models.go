package models

// SystemReport aggregates all collected telemetry data metrics
type SystemReport struct {
	// OS Metrics
	OSType      string  `json:"os_type"`
	Goroutines  int     `json:"goroutines"`
	CPUCores    int     `json:"cpu_cores"`
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	DiskUsage   float64 `json:"disk_usage"`

	// Docker Metrics
	DockerRunning int `json:"docker_running"`
	DockerStopped int `json:"docker_stopped"`

	// Kubernetes Metrics
	K8sNodes int `json:"k8s_nodes"`
	K8sPods  int `json:"k8s_pods"`

	// Overall Analysis Verdict
	Status string `json:"status"`
}