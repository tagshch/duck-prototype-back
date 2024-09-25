package controllers

import (
	"fmt"
	"net/http"
	"shell-backend/pkg/models"
)

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserModel

	err := c.UserStorage.CreateUser(user)

	if err != nil {
		SendJsonFailure(w, err.Error())
		return
	}

	SendJson(w, JsonResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("user '%s' is created", user.Name),
		Result:  nil,
	})
}

func (c *Controller) RemoveUser(w http.ResponseWriter, r *http.Request) {
	var userId string

	err := c.UserStorage.RemoveUser(userId)

	if err != nil {
		SendJsonFailure(w, err.Error())
		return
	}

	SendJson(w, JsonResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("user with id '%s' is removed", userId),
		Result:  nil,
	})
}

func (c *Controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.UserStorage.GetUsers()

	if err != nil {
		SendJsonFailure(w, err.Error())
		return
	}

	SendJson(w, JsonResponse{
		Status:  http.StatusOK,
		Message: "success",
		Result:  users,
	})
}

func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	var userId string

	user, err := c.UserStorage.GetUser(userId)

	if err != nil {
		SendJsonFailure(w, err.Error())
		return
	}

	SendJson(w, JsonResponse{
		Status:  http.StatusOK,
		Message: "success",
		Result:  user,
	})
}
