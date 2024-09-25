package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shell-backend/pkg/models"
)

func (c *Controller) CreateFile(w http.ResponseWriter, r *http.Request) {
	var file models.FileModel

	err := json.NewDecoder(r.Body).Decode(&file)

	if err != nil {
		SendJsonFailure(w, err.Error())
		return
	}

	err = c.FileStorage.CreateFile(file)

	if err != nil {
		SendJsonFailure(w, err.Error())
		return
	}

	SendJson(w, JsonResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("file '%s' is created", file.Table),
		Result:  nil,
	})
}

func (c *Controller) RemoveFile(w http.ResponseWriter, r *http.Request) {
	var file models.FileModel

	err := json.NewDecoder(r.Body).Decode(&file)

	if err != nil {
		SendJsonFailure(w, err.Error())
		return
	}

	err = c.FileStorage.RemoveFile(file.Id)

	if err != nil {
		SendJsonFailure(w, err.Error())
		return
	}

	SendJson(w, JsonResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("file with id '%d' is removed", file.Id),
		Result:  nil,
	})
}

func (c *Controller) GetFiles(w http.ResponseWriter, r *http.Request) {
	files, err := c.FileStorage.GetFiles()

	if err != nil {
		SendJsonFailure(w, err.Error())
		return
	}

	SendJson(w, JsonResponse{
		Status:  http.StatusOK,
		Message: "success",
		Result:  files,
	})
}
