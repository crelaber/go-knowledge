package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HTTPGet(uri string) ([]byte, error) {
	return HTTPContext(context.Background(), uri)
}

func HTTPContext(ctx context.Context, uri string) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error: uri=%s, statusCode=%v", uri, response.Status)
	}
	return ioutil.ReadAll(response.Body)
}

func HTTPPost(uri string, data string) ([]byte, error) {
	return HTTPPostWithContext(context.Background(), uri, data)
}

func HTTPPostWithContext(ctx context.Context, uri string, data string) ([]byte, error) {
	body := bytes.NewBuffer([]byte(data))
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post err : url=%s, statusCode=%v", uri, response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}

func PostJSON(url string, obj interface{}) ([]byte, error) {
	jsonBuf := new(bytes.Buffer)
	enc := json.NewEncoder(jsonBuf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(obj)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json;charset=utf-8", jsonBuf)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post err : url = %v , statusCode = %v", url, response.Status)
	}
	return ioutil.ReadAll(response.Body)
}
