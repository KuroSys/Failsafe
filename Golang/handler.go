package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"log"

	"github.com/go-resty/resty/v2"
)

type APIResponse struct {
	Current float32 `json:"Current"`
	Old     float32 `json:"Old"`
}

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
	if apiResponse.Current > apiResponse.Old {
		fmt.Println("[!] Found Newer version")
		fmt.Println("[i] Downloading Newest..")
		// Download Newest
		cmd := exec.Command("wget", "https://kurosys.bio/release/test")
		stdout, err := cmd.Output()
	
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(stdout))
		
	} else if (apiResponse.Current == apiResponse.Old) {
		fmt.Println("[+] You are on the latest version")
	}

}