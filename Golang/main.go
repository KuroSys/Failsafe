package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Monitor struct {
	Monitored []string `json:"Monitored"`
	Interval  int      `json:"Interval"`
	Screener  []string `json:"Screener"`
}

func main() {
	checkup()
	filePath := "settings.json"

	currentTime := time.Now()

	formattedTime := currentTime.Format("02.01. | 15:04:05")
	fmt.Printf("Current time: %s\n", formattedTime)

	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var monitor Monitor
	err = json.Unmarshal(jsonData, &monitor)
	if err != nil {
		log.Fatalf("Failed to parse JSON File: %v", err)
	}

	// Starte den Überwachungsprozess in einer Go-Routine
	go monitorScreens(monitor)
	select {}

	fmt.Printf("Interval: %d\n", monitor.Interval)
	fmt.Printf("Monitoring: %v\n", monitor.Monitored)
}

func executeCommand(cmd string) error {
	err := exec.Command("sh", "-c", cmd).Run()
	if err != nil {
		return err
	}
	return nil
}

func countdownTimer(interval int) {
	for i := interval; i > 0; i-- {
		fmt.Printf("\rChecking again in %d", i)
		time.Sleep(time.Second)
	}
}

func monitorScreens(monitor Monitor) {
	logFile, err := os.Create("Monitor.log")
	if err != nil {
		fmt.Println("[ERROR] Failed To Create Logfile!!")
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)

	for {
		allScreensRunning := true

		for _, serviceName := range monitor.Monitored {
			cmd := fmt.Sprintf("screen -list | grep -q %s", serviceName)
			err := executeCommand(cmd)
			if err != nil {
				allScreensRunning = false
				fmt.Printf("[ERROR] %s Is Not Active\n", serviceName)
				fmt.Printf("[INFO] Dead Screen Logged! -> Monitor.log\n")

				Logged := fmt.Sprintf("[ERROR] Screen %s is dead!", serviceName)
				logger.Printf(Logged)

				startScreener(serviceName, monitor.Screener)
			} else {
				fmt.Printf("[SUCCESS] %s Is Up And Running!\n", serviceName)
			}
		}

		if allScreensRunning {
			fmt.Println("[SUCCESS] All Screen Seems To Be Running\n")
		}
		countdownTimer(monitor.Interval)
		time.Sleep(time.Second * time.Duration(monitor.Interval))
		banner()
	}
}

func startScreener(serviceName string, screenerCommands []string) {
	for _, screenerCmd := range screenerCommands {
		if strings.Contains(screenerCmd, serviceName) {
			err := exec.Command("sh", "-c", screenerCmd).Run()
			if err != nil {
				fmt.Printf("[PANIC] Error Starting Screener For %s: %v\n", serviceName, err)
			} else {
				fmt.Printf("[INFO] Screener Started For %s\n", serviceName)
			}
		}
	}
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
