package util

import (
	"encoding/json"
	"net/url"
	"io"
	"fmt"
	"strings"
)

func GetQueryParams(request interface{}) (urlEncoded string, err error) {
	jsonStr, err := json.Marshal(request)
	if err != nil {
		return
	}
	var result map[string]interface{}
	if err = json.Unmarshal(jsonStr, &result); err != nil {
		return
	}

	urlEncoder := url.Values{}
	for key, value := range result {
		if value != "" {
			urlEncoder.Add(key, value.(string))
		}
	}
	urlEncoded = urlEncoder.Encode()
	return
}

func GetBodyParams(request interface{}) (reader io.Reader, err error) {
	jsonStr, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(string(jsonStr)), nil
}
