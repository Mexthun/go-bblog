package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetRequestJsonParam(r *http.Request) (data map[string]interface{}, err error) {
	requestBody, err := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(requestBody, &data)

	if err != nil {
		return nil, err
	}
	return data, nil
}
