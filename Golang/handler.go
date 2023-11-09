package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gookit/color"
)

type APIResponse struct {
	Current float32 `json:"Current"`
	Old     float32 `json:"Old"`
}

type MonitorHandler struct {
	Monitored []string `json:"Monitored"`
	Interval  int      `json:"Interval"`
}

const Vers = 1.1

func checkup() {

	apiURL := "https://kurosys.bio/api/check.json"

	client := resty.New()

	response, err := client.R().
		Get(apiURL)

	if err != nil {
		log.Fatalf("[FATAL] Failed to Request API: %v", err)
	}

	var apiResponse APIResponse
	if err := json.Unmarshal(response.Body(), &apiResponse); err != nil {
		log.Fatalf("[FATAL] Failed to parase API: %v", err)
	}

	fmt.Println("Checking For Updates...")
	if apiResponse.Current > Vers {
		fmt.Println("[!] Found Newer version")
		fmt.Println("[i] Downloading Newest..")

		if err := os.Remove("handler"); err != nil {
			fmt.Println("[ERROR] Failed To Delete Old File")
		}

		cmd := exec.Command("wget", "https://kurosys.bio/release/handler", "-q")
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println("[x] Failed To Download New File", err)
			return
		}
		fmt.Println(string(stdout))
		cmd = exec.Command("chmod", "+x", "handler")
		fmt.Println("[+] Updated Successfully")
		time.Sleep(2 * time.Second)
		fmt.Println("[i] Restarting...")
		time.Sleep(1 * time.Second)
		cmd = exec.Command("./handler")
	} else if Vers == apiResponse.Current {
		fmt.Println("[+] You are on the latest version")
		time.Sleep(2 * time.Second)
	} else if Vers > apiResponse.Current {
		fmt.Println("You are on an unreleased Version? imao")
		os.Exit(0)
	}
	banner()
}

func banner() {
	clearConsole()
	banner := `
	██╗  ██╗██╗   ██╗██████╗  ██████╗ 
	██║ ██╔╝██║   ██║██╔══██╗██╔═══██╗
	█████╔╝ ██║   ██║██████╔╝██║   ██║
	██╔═██╗ ██║   ██║██╔══██╗██║   ██║
	██║  ██╗╚██████╔╝██║  ██║╚██████╔╝
	╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ 	
	sys(.bio)		
		`
	red := color.New(color.FgRed)
	red.Println(banner)
}
