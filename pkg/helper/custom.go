package helper

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"zclient/pkg/lib"

	"github.com/gojek/heimdall/v7/httpclient"
	jsoniter "github.com/json-iterator/go"
)

type Auth struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Message      string `json:"message"`
}

func CheckSession() (bool, error) {
	var data Auth

	file, err := os.Open(".session")
	if err != nil {
		return false, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	err = jsoniter.NewDecoder(file).Decode(&data)

	if data.Token == "" || data.RefreshToken == "" {
		return false, err
	}

	url := os.Getenv("API_URL") + "checktoken"
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := httpclient.NewClient(httpclient.WithHTTPClient(&lib.HttpClient{Client: http.Client{Transport: transport}}))
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return false, err
	}

	request.Header.Set("Authorization", "Bearer "+data.Token)
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)

	if err != nil {
		return false, err
	}

	body, _ := io.ReadAll(response.Body)
	if err := jsoniter.Unmarshal(body, &data); err != nil {
		return false, err
	}

	if data.Message != "Token Is Valid" {
		return false, err
	}

	return true, nil
}

func GetToken() *Auth {
	var data Auth
	file, err := os.Open(".session")
	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	err = jsoniter.NewDecoder(file).Decode(&data)

	if err != nil {
		fmt.Println(err)
	}

	return &data
}

func CreateSession(token string, refreshtoken string) {
	file, err := os.Create(".session")

	if err != nil {
		fmt.Print(err)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	tokenAuth := &Auth{
		Token:        token,
		RefreshToken: refreshtoken,
	}

	jsonValue, err := jsoniter.Marshal(tokenAuth)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(jsonValue)

}

func Logout() {
	err := os.Remove(".session")
	if err != nil {
		fmt.Println(err)
	}
}
