package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

type Monitor struct {
	Monitored []string `json:"Monitored"`
	Interval  int      `json:"Interval"`
}

func main() {

	checkup()

	filePath := "settings.json"
	sleepTime := monitor.Interval

	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var monitor Monitor
	err = json.Unmarshal(jsonData, &monitor)
	if err != nil {
		log.Fatalf("Failed to parse JSON File: %v", err)
	}

	fmt.Printf("Interval: %d\n", monitor.Interval)
	fmt.Printf("Monitoring: %v\n", monitor.Monitored)
	var bool cycle = true

	

	for _, serviceName := range monitor.Monitored {
		cmd := fmt.Sprintf("screen -list | grep -q %s", serviceName)
		err := executeCommand(cmd)
		if err != nil {
			fmt.Printf("[FAILED] Screen '%s' is not running | %v\n", serviceName, err)
			fmt.Println("[FAILED] Attempting to start screen")
		} else {
			fmt.Printf("[SUCCESS] Screen '%s' is active and running\n", serviceName)
			fmt.Printf("[SUCCESS] Check Completed Without Errors. Sleeping For %d Seconds\n", sleepTime)
		}
	}
}

func executeCommand(cmd string) error {
	err := exec.Command("sh", "-c", cmd).Run()
	if err != nil {
		return err
	}
	return nil
}
