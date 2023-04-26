package util

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"zclient/pkg/lib"

	"github.com/gojek/heimdall/v7/httpclient"
	jsoniter "github.com/json-iterator/go"
)

type ResponseProfile struct {
	Data struct {
		Username    string   `json:"username"`
		Apikey      string   `json:"apikey"`
		Email       string   `json:"email"`
		WhitelistIp []string `json:"whitelist_ip"`
		BlacklistIp []string `json:"blacklist_ip"`
	} `json:"data"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type EditProfile struct {
	Email           string   `json:"email"`
	Username        string   `json:"username"`
	Password        string   `json:"password"`
	ConfirmPassword string   `json:"confirm_password"`
	WhitelistIp     []string `json:"whitelist_ip"`
	BlacklistIp     []string `json:"blacklist_ip"`
}

func (S *ResponseProfile) GetProfile(token string, refreshtoken string) (*ResponseProfile, error) {
	var data *ResponseProfile
	url := os.Getenv("API_URL") + "profile/"
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := httpclient.NewClient(httpclient.WithHTTPClient(&lib.HttpClient{Client: http.Client{Transport: transport}}))
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)
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

func (P *EditProfile) UpdateProfile(token string, refreshtoken string) (*ResponseProfile, error) {
	values := make(map[string]string)
	var data *ResponseProfile
	url := os.Getenv("API_URL") + "profile/update"
	if P.Email != "" {
		// values := map[string]string{
		// 	"email":            P.Email,
		// 	"username":         P.Username,
		// 	"password":         P.Password,
		// 	"confirm_password": P.ConfirmPassword,
		// 	//"whitelist_ip":     P.WhitelistIp[0],
		// 	//"blacklist_ip":     P.BlacklistIp[0],
		// }
		values["email"] = P.Email
	}
	if P.Username != "" {
		values["username"] = P.Username
	}
	if P.Password != "" && P.ConfirmPassword != "" {
		values["password"] = P.Password
		values["confirm_password"] = P.ConfirmPassword
	}
	jsonValue, _ := json.Marshal(values)
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := httpclient.NewClient(httpclient.WithHTTPClient(&lib.HttpClient{Client: http.Client{Transport: transport}}))
	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer([]byte(jsonValue)))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)
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

func (P *EditProfile) AddIp(token string, refreshtoken string) (*ResponseProfile, error) {
	var data *ResponseProfile
	url := os.Getenv("API_URL") + "profile/update"
	values := map[string][]string{
		"whitelist_ip": P.WhitelistIp,
		"blacklist_ip": P.BlacklistIp,
	}
	jsonValue, _ := json.Marshal(values)
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := httpclient.NewClient(httpclient.WithHTTPClient(&lib.HttpClient{Client: http.Client{Transport: transport}}))
	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer([]byte(jsonValue)))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)
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
