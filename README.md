# DevOps Monitor CLI

A lightweight, high-performance, cross-platform command-line utility built with **Go** that aggregates host telemetry, Docker container status, and Kubernetes cluster metrics into a unified, human-readable infrastructure health report.

---

## 🚀 Features

### Real-Time Host Diagnostics

Collects live system telemetry using cross-platform system APIs powered by `gopsutil`, including:

* CPU utilization and load statistics
* Memory consumption metrics
* Disk usage and capacity information
* Active process and thread insights

### Parallel Metrics Collection

Leverages Go concurrency primitives for maximum performance:

* Goroutines
* `sync.WaitGroup`
* `sync.Mutex`

All subsystem checks execute concurrently to minimize scan latency.

### Docker & Kubernetes Monitoring

Integrates directly with local orchestration environments:

* Docker daemon status and container states
* Kubernetes cluster node and pod metrics
* Local `kubectl` context inspection

### Intelligent Health Analysis

Converts raw telemetry into actionable infrastructure health states:

| Status   | Description                                   |
| -------- | --------------------------------------------- |
| HEALTHY  | System operating within acceptable thresholds |
| WARNING  | Resource usage approaching limits             |
| CRITICAL | Resource usage exceeds safe thresholds        |

---

# 🏗 Architecture

The application follows a modular and thread-safe workflow to ensure efficient metric collection and processing.

```text
                  +-------------------------+
                  |    devops-monitor CLI   |
                  +------------+------------+
                               |
            +------------------+------------------+
            |                  |                  |
            v                  v                  v
    +---------------+  +---------------+  +---------------+
    |   OS Engine   |  | Docker Daemon |  | Kubernetes API|
    |  (Goroutine)  |  |  (Goroutine)  |  |  (Goroutine)  |
    +-------+-------+  +-------+-------+  +-------+-------+
            |                  |                  |
            +------------------+------------------+
                               |
                               | [sync.Mutex]
                               v
                  +-------------------------+
                  |  Metrics Cache Struct   |
                  +------------+------------+
                               |
                               v
                  +-------------------------+
                  |  Health Analyzer Engine |
                  +------------+------------+
                               |
                               v
                  +-------------------------+
                  | Terminal Report Output  |
                  +-------------------------+
```

---

## 🔄 Workflow

### 1. Trigger

The operator executes the CLI with a scan command.

### 2. Collection

Independent goroutines gather:

* Host operating system metrics
* Docker container information
* Kubernetes cluster statistics

### 3. Aggregation

Collected metrics are synchronized and stored safely using mutex-protected structures.

### 4. Analysis

The health engine evaluates collected data against predefined thresholds.

### 5. Reporting

A consolidated infrastructure report is rendered to the terminal.

---

# 📂 Project Structure

```text
devops-monitor/
├── .github/
│   └── workflows/
│       └── release.yml      # GitHub Actions CI/CD pipeline
├── bin/                     # Compiled executables
├── internal/
│   ├── collector/
│   │   ├── orchestrator.go  # Docker & Kubernetes collectors
│   │   └── os.go            # Host system telemetry collector
│   ├── engine/
│   │   └── scanner.go       # Parallel execution and health scoring
│   ├── models/
│   │   └── models.go        # Shared telemetry models
│   └── report/
│       └── renderer.go      # Terminal report renderer
├── build.sh                 # Build automation script
├── go.mod                   # Dependency definitions
├── go.sum                   # Dependency checksums
└── main.go                  # Application entry point
```

---

# 📦 Installation

## Windows

### Option A — Winget (Recommended)

```powershell
winget install devops-monitor
```

### Option B — Manual Installation

1. Download the latest Windows binary from the project's Releases page.

2. Rename the executable:

```powershell
Rename-Item .\devops-monitor-windows-amd64.exe .\devops-monitor.exe
```

3. Move the binary to a preferred directory.

4. Add the directory to your system `PATH`.

---

## Linux & macOS

### Option A — Homebrew

```bash
brew tap srinivas-batthula/tap
brew install devops-monitor
```

### Option B — SnapCraft

```bash
sudo snap install devops-monitor --classic
```

### Option C — Manual Installation

#### Linux (AMD64)

```bash
curl -L -o devops-monitor \
https://github.com/srinivas-batthula/devops-monitor/releases/download/v1.0.2/devops-monitor-linux-amd64

chmod +x devops-monitor

sudo mv devops-monitor /usr/local/bin/
```

#### macOS (Apple Silicon)

```bash
curl -L -o devops-monitor \
https://github.com/srinivas-batthula/devops-monitor/releases/download/v1.0.2/devops-monitor-darwin-arm64

chmod +x devops-monitor

sudo mv devops-monitor /usr/local/bin/
```

---

# ▶ Usage

Run a complete infrastructure scan:

```bash
devops-monitor scan
```

---

# 📊 Sample Output

```text
========================================
         DEVOPS MONITOR REPORT
========================================

OS Architecture   : windows
Active Goroutines : 4
Logical CPU Cores : 8
CPU Utilization   : 24.52%
Memory Utilization: 58.11%
Disk Allocation   : 47.30%

----------------------------------------
Docker Container State:
  Running         : 3
  Stopped         : 1

----------------------------------------
Kubernetes Cluster State:
  Active Nodes    : 1
  Monitored Pods  : 14

----------------------------------------
OVERALL ENGINE HEALTH STATUS: HEALTHY

========================================
```

---

# ⚡ Technology Stack

* **Go**
* **gopsutil**
* **Docker CLI**
* **kubectl**
* **GitHub Actions**
* **Goroutines**
* **sync.WaitGroup**
* **sync.Mutex**

---

# 🔐 Requirements

Ensure the following tools are available if Docker and Kubernetes monitoring are enabled:

* Docker Engine / Docker Desktop
* Kubernetes CLI (`kubectl`)
* Valid Kubernetes context configuration

---

# 📜 License

Licensed under the MIT License.

---

# 🤝 Contributing

Contributions, bug reports, and feature requests are welcome.

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Open a Pull Request

---

Built with ❤️ using Go.
