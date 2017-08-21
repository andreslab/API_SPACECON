package main

import (
	"encoding/json"
	"net/http"
)

var responseDataCheckService = make(map[string]CheckServicesController)

type CheckServicesController struct {
	ID    string `json:"ID"`
	STATE string `json:"STATE"`
	ERROR string `json:"ERROR"`
}

func NewCheckServicesControllerEmpty() *CheckServicesController {
	return &CheckServicesController{
		ID:    "0",
		STATE: "0",
		ERROR: "0",
	}
}

func NewCheckServicesController(id string, state string, _error string) *CheckServicesController {
	return &CheckServicesController{
		ID:    id,
		STATE: state,
		ERROR: _error,
	}
}

func CheckServicesRequestGet(w http.ResponseWriter, r *http.Request) {

	dataToSend := NewResponseController("0", "1", "0", "200")
	//header
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(dataToSend)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
