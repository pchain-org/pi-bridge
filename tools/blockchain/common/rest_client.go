package common

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type RestClient struct {
	Addr       string
	restClient *http.Client
}

func NewRestClient() *RestClient {
	return &RestClient{
		restClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   5,
				DisableKeepAlives:     false,
				IdleConnTimeout:       time.Second * 300,
				ResponseHeaderTimeout: time.Second * 300,
				TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: time.Second * 300,
		},
	}
}

func (rc *RestClient) SetAddr(addr string) *RestClient {
	rc.Addr = addr
	return rc
}

func (rc *RestClient) SetRestClient(restClient *http.Client) *RestClient {
	rc.restClient = restClient
	return rc
}

func (rc *RestClient) SendRestRequest(addr string, data []byte) ([]byte, error) {
	resp, err := rc.restClient.Post(addr, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, fmt.Errorf("http post request:%s error:%s", data, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read rest response body error:%s", err)
	}
	return body, nil
}
