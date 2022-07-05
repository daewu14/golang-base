package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HttpService() NetClient {
	return NetClient{}
}

type NetClient struct {
	url     string
	method  string
	data    map[string]string
	headers []headers
	body    interface{}
}

type headers struct {
	key   string
	value string
}

func (cl NetClient) Get() NetClient {
	cl.method = http.MethodGet
	return cl
}

func (cl NetClient) Put() NetClient {
	cl.method = http.MethodPut
	return cl
}

func (cl NetClient) Delete() NetClient {
	cl.method = http.MethodDelete
	return cl
}

// Post format must you use
//
// Post().
// Url("url").
// AddHeader("XX", "xx").
// Params(data).
// Call()
//
// if the method always return error please use http.PostForm()
func (cl NetClient) Post() NetClient {
	cl.method = http.MethodPost
	return cl
}

func (c NetClient) Url(url string) NetClient {
	c.url = url
	return c
}

func (c NetClient) Params(data map[string]string) NetClient {
	c.data = data
	return c
}

func (c NetClient) Bodys(data interface{}) NetClient {
	c.body = data
	return c
}

func (c NetClient) AddHeader(key string, value string) NetClient {
	c.headers = append(c.headers, headers{
		key:   key,
		value: value,
	})
	return c
}

func (c NetClient) Call() (*http.Response, error) {
	jsonData, _ := json.Marshal(c.data)
	println(c.method+" url :", c.url)
	println("parameters :", string(jsonData))

	var body io.Reader
	fmt.Println(body)
	var cek *bytes.Buffer = nil

	switch c.method {
	case http.MethodGet:
		body = nil
		break
	case http.MethodPost:
		cek = new(bytes.Buffer)
		json.NewEncoder(cek).Encode(c.body)
		fmt.Println("cookk", cek)

	case http.MethodPut:
	case http.MethodDelete:
		body = bytes.NewBuffer(jsonData)
		break
	}

	request, err := http.NewRequest(c.method, c.url, cek)
	for _, val := range c.headers {
		request.Header.Add(val.key, val.value)
	}

	if c.method == http.MethodGet {
		q := request.URL.Query()
		for s := range c.data {
			q.Add(s, c.data[s])
		}
		request.URL.RawQuery = q.Encode()
	}

	if err != nil {
		return &http.Response{}, err
	}
	client := http.Client{}
	return client.Do(request)
}
