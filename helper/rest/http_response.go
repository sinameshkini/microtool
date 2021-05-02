package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Code    int        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

func ReturnResponse(w http.ResponseWriter, status string, code int, message string, data, meta interface{}) {
	var (
		err error
		resp = Response{
			Status:  status,
			Code:    code,
			Message: message,
			Data:    data,
			Meta:    meta,
		}
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
	}
}

