package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var client *http.Client

func init()  {
	client = &http.Client{}
}


func Post(url string, body io.Reader, headers map[string]string) (map[string]interface{}, error) {
	request, _ := http.NewRequest("POST", url, body)

	if headers != nil {
		for key, value := range headers {
			request.Header.Set(key, value)
		}
	}
	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if result, err = toMap(res.Body); err != nil {
		return nil, err
	}
	return result, nil
}

func Get(url string, headers map[string]string) (map[string]interface{}, error)  {
	request, _ := http.NewRequest("GET", url, nil)
	if headers != nil {
		for key, value := range headers {
			request.Header.Set(key, value)
		}
	}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var result map[string]interface{}
	if result, err = toMap(res.Body); err != nil {
		return nil, err
	}
	return result, nil
}

func toMap(body io.Reader) (map[string]interface{}, error)  {
	var result map[string]interface{}
	data, _ := ioutil.ReadAll(body)
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
