package pteroclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var client = http.Client{
	Timeout: 30 * time.Second,
}

var authorization = fmt.Sprintf("Bearer %s", os.Getenv("PTERO_API_KEY"))

var requestURL = os.Getenv("PTERO_API_URL")

func NewPteroRequest(method string, path string) (*http.Response, error) {
	req, err := http.NewRequest(method, requestURL+path, nil)
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Accept", "Application/vnt.pterodactyl.v1+json")
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func Servers() (*ServerList, error) {
	res, err := NewPteroRequest(http.MethodGet, "client")
	if err != nil {
		return nil, err
	}
	list := &ServerList{}
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, list)
	return list, err
}
