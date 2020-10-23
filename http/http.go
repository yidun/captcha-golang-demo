package http

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var httpClient = &http.Client{}

type HeaderOption struct {
	Name  string
	Value string
}

type QueryParameter struct {
	Key   string
	Value interface{}
}

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient = &http.Client{
		Transport: tr,
		Timeout:   time.Duration(30) * time.Second,
	}
}

func PostForm(url string, params map[string]string,) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(convertToQueryParams(params)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := httpClient.Do(req)
	defer func() {
		if resp != nil {
			if e := resp.Body.Close(); e != nil {
				fmt.Println(e)
			}
		}
	}()
	return responseHandle(resp, err)
}

func responseHandle(resp *http.Response, err error) (string, error) {
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	respBody := string(b)
	return respBody, nil
}

func convertToQueryParams(params map[string]string) string {
	if params == nil || len(params) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for k, v := range params {
		buffer.WriteString(fmt.Sprintf("%s=%v&", k, v))
	}
	buffer.Truncate(buffer.Len() - 1)
	return buffer.String()
}
