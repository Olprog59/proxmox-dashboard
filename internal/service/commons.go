package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Olprog59/dashboard-proxmox/internal/config"
)

// type Client struct {
// 	baseUrl     string
// 	httpClient  *http.Client
// 	tokenId     string
// 	tokenSecret string
// }
//
// func NewClientProxmox(baseUrl, tokenId, tokenSecret string) *Client {
// 	return &Client{
// 		baseUrl:     baseUrl,
// 		httpClient:  &http.Client{Timeout: 30 * time.Second},
// 		tokenId:     tokenId,
// 		tokenSecret: tokenSecret,
// 	}
// }

func DoRequest[T any](method, path string, c config.ClusterConfig, data any) ([]T, error) {
	var reqBody io.Reader
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	fullURL := fmt.Sprintf("%s/api2/json/%s", c.APIURL, strings.TrimPrefix(path, "/"))
	req, err := http.NewRequest(method, fullURL, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "PVEAPIToken="+c.SecretID+"="+c.SecretToken)

	if data != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("erreur API: %s, code: %d", body, resp.StatusCode)
	}

	var result struct {
		Data []T `json:"data"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}
