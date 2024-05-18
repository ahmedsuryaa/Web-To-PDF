package main

import (
	"encoding/json"
	"fmt"
	"github.com/extism/go-pdk"
)

var API_KEY string = "urlbox_api_key"

type renderResponse struct {
	RenderUrl string `json:"renderUrl"`
	Size      int64  `json:"size"`
}

//export run
func run() int32 {

	WebsiteURL := pdk.InputString()
	if WebsiteURL == "" {
		pdk.SetError(fmt.Errorf("no URL provided"))
		return 1
	}
	
	payload := fmt.Sprintf(`{
		"url": "%s",
		"format": "pdf",
		"full_page": true
	}`, WebsiteURL)

	url := "https://api.urlbox.io/v1/render/sync"

	req := pdk.NewHTTPRequest(pdk.MethodPost, url)
	req.SetHeader("Content-Type", "application/json")
	req.SetHeader("Authorization", "Bearer " + API_KEY)
	req.SetBody([]byte(payload))
	
	res := req.Send()

	if res.Status() != 200 {
		pdk.SetError(fmt.Errorf("failed to render URL: %v", res.Status()))
		return 1
	}

	var renderRes renderResponse
	if err := json.Unmarshal(res.Body(), &renderRes); err != nil {
		pdk.SetError(err)
		return 1
	}

	pdk.OutputString("PDF URL: " + renderRes.RenderUrl + "\n")

	return 0
}

func main() {}