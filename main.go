package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

func main() {
	// Parse the command-line arguments
	flag.Parse()

	// Retrieve the port number from the command-line arguments
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: freeport <port_number>")
		os.Exit(1) // Use os.Exit(1) to exit the program with an error status
	}

	port := args[0]

	// Check the operating system
	switch runtime.GOOS {
	case "darwin", "linux":
		// For Linux and macOS, use lsof and kill commands
		processes := findProcessesUsingPortUnix(port)
		if len(processes) == 0 {
			fmt.Printf("No processes using port %s\n", port)
			return
		}
		// Map to track terminated processes
		terminated := make(map[string]bool)
		var mutex sync.Mutex

		for {
			// Terminate processes
			for _, process := range processes {
				mutex.Lock()
				if terminated[process.PID] {
					mutex.Unlock()
					continue
				}
				mutex.Unlock()

				err := terminateProcessUnix(process.PID)
				if err != nil {
					fmt.Printf("Error terminating process %s (PID: %s): %v\n", process.Name, process.PID, err)
				} else {
					fmt.Printf("Process %s (PID: %s) terminated\n", process.Name, process.PID)
				}

				mutex.Lock()
				terminated[process.PID] = true
				mutex.Unlock()
			}
			// Check again if any processes are still running
			processes = findProcessesUsingPortUnix(port)
			if len(processes) == 0 {
				fmt.Printf("All processes using port %s terminated\n", port)
				break
			}
		}
	case "windows":
		// For Windows, use tasklist to create a map of PIDs and process names
		pidMap := createProcessMap()

		// For Windows, use netstat and taskkill commands
		processes := findProcessesUsingPortWindows(port, pidMap)
		if len(processes) == 0 {
			fmt.Printf("No processes using port %s\n", port)
			return
		}

		// Map to track terminated processes
		terminated := make(map[string]bool)
		var mutex sync.Mutex

		for {
			for _, process := range processes {
				mutex.Lock()
				if terminated[process.PID] {
					mutex.Unlock()
					continue
				}
				mutex.Unlock()
				
				err := terminateProcessWindows(process.PID)
				if err != nil {
					fmt.Printf("Error terminating process %s (PID: %s): %v\n", process.Name, process.PID, err)
				} else {
					fmt.Printf("Process %s (PID: %s) terminated\n", process.Name, process.PID)
				}

				mutex.Lock()
				terminated[process.PID] = true
				mutex.Unlock()
			}
			// Check again if any processes are still running
			processes = findProcessesUsingPortWindows(port, pidMap)
			if len(processes) == 0 {
				fmt.Printf("All processes using port %s terminated\n", port)
				break
			}
		}
	default:
		fmt.Printf("Unsupported operating system: %s\n", runtime.GOOS)
		os.Exit(1)
	}
}

// ProcessInfo represents information about a process
type ProcessInfo struct {
	Name string
	PID  string
}

// findProcessesUsingPortUnix finds processes using the specified port on Unix-like systems (Linux, macOS)
func findProcessesUsingPortUnix(port string) []ProcessInfo {
	cmd := exec.Command("lsof", "-i", ":"+port)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(err.Error(), "exit status 1") {
            // No processes found, return empty slice
            return []ProcessInfo{}
        }
		fmt.Printf("Error running lsof command: %v\n", err)
		os.Exit(1)
	}

	lines := strings.Split(string(output), "\n")
	var processes []ProcessInfo
	for _, line := range lines[1:] { // Skip the header line
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			processes = append(processes, ProcessInfo{Name: fields[0], PID: fields[1]})
		}
	}
	return processes
}

// terminateProcessUnix terminates the specified process on Unix-like systems (Linux, macOS)
func terminateProcessUnix(pid string) error {
	cmd := exec.Command("kill", pid)
	err := cmd.Run()
	return err
}

func createProcessMap() map[string]string {
	// Run tasklist command to get process names
    cmdTasklist := exec.Command("tasklist")
    outputTasklist, err := cmdTasklist.CombinedOutput()
    if err != nil {
        fmt.Printf("Error running tasklist command: %v\n", err)
        os.Exit(1)
    }

	// Parse tasklist output to get PIDs and process names
	tasklistLines := strings.Split(string(outputTasklist), "\n")
	pidMap := make(map[string]string)
	for _, line := range tasklistLines[3:] { // Skip header lines
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			pidMap[fields[1]] = fields[0]
		}
	}
	return pidMap
}

// findProcessesUsingPortWindows finds processes using the specified port on Windows
func findProcessesUsingPortWindows(port string, pidMap map[string]string) []ProcessInfo {
    // Run netstat command to get active connections
    cmdNetstat := exec.Command("netstat", "-aon")
    outputNetstat, err := cmdNetstat.CombinedOutput()
    if err != nil {
        fmt.Printf("Error running netstat command: %v\n", err)
        os.Exit(1)
    }

    lines := strings.Split(string(outputNetstat), "\n")

	var processes []ProcessInfo
    for _, line := range lines[4:] { // Skip header lines
        fields := strings.Fields(line)
        if len(fields) >= 4 && strings.HasSuffix(fields[1], ":"+port) {
			pid := fields[4]
			name := pidMap[pid]
            processes = append(processes, ProcessInfo{Name: name, PID: pid})
        }
    }
    return processes
}

// terminateProcessWindows terminates the specified process on Windows
func terminateProcessWindows(pid string) error {
	cmd := exec.Command("taskkill", "/F", "/PID", pid)
	err := cmd.Run()
	return err
}
