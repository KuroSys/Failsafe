package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type APIResponse struct {
	Current float32 `json:"Current"`
	Old     float32 `json:"Old"`
}

func main() {

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

	// fmt.Println("Made by KuroSys(.bio)\n")
	// fmt.Printf("v%.1f\n", apiResponse.Version)

	fmt.Println("Checking For Updates...")
	if apiResponse.Current > apiResponse.Old {
		fmt.Println("[!] Found Newer version")
		fmt.Println("[i] Downloading Newest")
		resp, err := http.Get("https://kurosys.bio/release/failsafe")
		defer out.Close()
	}
}
