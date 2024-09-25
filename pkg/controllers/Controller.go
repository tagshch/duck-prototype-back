package controllers

import (
	"encoding/json"
	"net/http"
	"shell-backend/pkg/storages"
)

type Controller struct {
	UserStorage *storages.UserStorage
	FileStorage *storages.FileStorage
}

type JsonResponse struct {
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

func SendJson(w http.ResponseWriter, data JsonResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.Status)
	json.NewEncoder(w).Encode(data)
}

func SendJsonSuccess(w http.ResponseWriter, data interface{}) {
	SendJson(w, JsonResponse{
		Message: "success",
		Status:  http.StatusOK,
		Result:  data,
	})
}

func SendJsonSuccessWithMessage(w http.ResponseWriter, message string, data interface{}) {
	SendJson(w, JsonResponse{
		Message: message,
		Status:  http.StatusOK,
		Result:  data,
	})
}

func SendJsonFailure(w http.ResponseWriter, message string) {
	SendJson(w, JsonResponse{
		Message: message,
		Status:  http.StatusBadRequest,
	})
}
