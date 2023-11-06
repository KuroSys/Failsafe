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

func spawner() {
	api()
}

func main() {

	spawner()

	filePath := "settings.json"

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

	for _, serviceName := range monitor.Monitored {
		cmd := fmt.Sprintf("screen -list | grep -q %s", serviceName)
		err := executeCommand(cmd)
		if err != nil {
			fmt.Printf("[FAILED] Screen '%s' is not running | %v\n", serviceName, err)
		} else {
			fmt.Printf("[SUCCESS] Screen '%s' is active and running\n", serviceName)
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
