package common

import (
	"encoding/json"
	"me-bblog/models"
	"net/http"
)

func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = 999
	result.Error = err.Error()
	jsondata, errJson := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	if errJson != nil {
		panic(err)
	}
	w.Write(jsondata)
}
func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Data = data
	jsondata, err := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	w.Write(jsondata)
}
