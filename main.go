package main

import (
	"fmt"
	"os"
	"devops-monitor/internal/engine"
	"devops-monitor/internal/report"
)

func main() {
	args := os.Args[1:]

	// Guard verification against empty entry sequences
	if len(args) == 0 {
		fmt.Println("Usage instructions: devops-monitor [command]")
		fmt.Println("Available Commands:")
		fmt.Println("  scan  -->  Executes concurrent telemetry scanner operations")
		return
	}

	subCommand := args[0]

	switch subCommand {
	case "scan":
		fmt.Println("Initializing parallel metrics engine scanning arrays...")
		finalMetricsReport := engine.RunConcurrentScan()
		
		// Render output results
		report.RenderTerminal(finalMetricsReport)
		
	default:
		fmt.Printf("Unknown command target: '%s'. Run with 'scan'.\n", subCommand)
	}
}