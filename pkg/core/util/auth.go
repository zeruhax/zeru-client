package util

import (
	"bytes"
	"crypto/tls"
	"github.com/gojek/heimdall/v7/httpclient"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"os"
	"zclient/pkg/lib"
)

type Auth struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseAuth struct {
	Message      string `json:"message"`
	Error        string `json:"error"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (A *Auth) Login() (*ResponseAuth, error) {
	var data *ResponseAuth
	url := os.Getenv("API_URL") + "login"
	values := map[string]string{"email": A.Email, "password": A.Password}
	jsonValue, _ := jsoniter.Marshal(values)
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := httpclient.NewClient(httpclient.WithHTTPClient(&lib.HttpClient{Client: http.Client{Transport: transport}}))
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonValue)))
	if err != nil {
		return nil, err
	}
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

func (A *Auth) Register() (*ResponseAuth, error) {
	var data *ResponseAuth
	url := os.Getenv("API_URL") + "register"
	values := map[string]string{"email": A.Email, "username": A.Username, "password": A.Password}
	jsonValue, _ := jsoniter.Marshal(values)
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := httpclient.NewClient(httpclient.WithHTTPClient(&lib.HttpClient{Client: http.Client{Transport: transport}}))
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonValue)))
	if err != nil {
		return nil, err
	}
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
