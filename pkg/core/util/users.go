package util

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"
	"zclient/pkg/lib"

	"github.com/gojek/heimdall/v7/httpclient"
	jsoniter "github.com/json-iterator/go"
)

type ResponseUsers struct {
	Result struct {
		Data []struct {
			Username  string `json:"username"`
			CreatedAt string `json:"created_at"`
			UpdateAt  string `json:"update_at"`
		} `json:"data"`
	} `json:"result"`
}

func GetUsers(token string) (*ResponseUsers, error) {
	var data *ResponseUsers
	url := os.Getenv("API_URL") + "users"
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := httpclient.NewClient(httpclient.WithHTTPClient(&lib.HttpClient{Client: http.Client{Transport: transport}}))
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(response.Body)
	if err := jsoniter.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}
