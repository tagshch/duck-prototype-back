package main

import (
	"log"
	"net/http"
	"os"

	"shell-backend/pkg/controllers"
	"shell-backend/pkg/middlewares"
	"shell-backend/pkg/models"
	"shell-backend/pkg/storages"
)

func main() {
	os.Remove("./testdb.db")

	storage, err := storages.NewSqlLiteStorage("./testdb.db")

	if err != nil {
		log.Fatal("cannot create db connection")
	}

	defer storage.DB.Close()

	userStorage := &storages.UserStorage{Storage: storage}

	err = userStorage.CreateTable()

	if err != nil {
		log.Fatal(err.Error())
	}

	fileStorage := &storages.FileStorage{Storage: storage}

	err = fileStorage.CreateTable()

	if err != nil {
		log.Fatal(err.Error())
	}

	// add default files
	err = fileStorage.CreateFile(models.FileModel{
		Table:  "2023_4",
		Url:    "https://storage.yandexcloud.net/duckdbshell/2023_4.json",
		UserId: "1",
	})

	if err != nil {
		log.Fatal("default values: ", err.Error())
	}

	err = fileStorage.CreateFile(models.FileModel{
		Table:  "2024_1",
		Url:    "https://storage.yandexcloud.net/duckdbshell/2024_1.json",
		UserId: "1",
	})

	if err != nil {
		log.Fatal("default values: ", err.Error())
	}
	// end add default files

	mux := http.NewServeMux()

	ctrl := controllers.Controller{
		UserStorage: userStorage,
		FileStorage: fileStorage,
	}

	mux.Handle("/api/user/add", http.HandlerFunc(ctrl.CreateUser))
	mux.Handle("/api/user/remove", http.HandlerFunc(ctrl.RemoveUser))
	mux.Handle("/api/user/list", http.HandlerFunc(ctrl.GetUsers))
	mux.Handle("/api/user/one", http.HandlerFunc(ctrl.GetUser))

	mux.Handle("/api/file/add", http.HandlerFunc(ctrl.CreateFile))
	mux.Handle("/api/file/remove", http.HandlerFunc(ctrl.RemoveFile))
	mux.Handle("/api/file/list", http.HandlerFunc(ctrl.GetFiles))

	handler := middlewares.Logger(mux)
	handler = middlewares.Cors(handler, "*")

	log.Println("Starting server!")

	err = http.ListenAndServe(":8081", handler)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server started!")

	os.Exit(0)
}
