package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"sihmba_server.go/handler"
	"sihmba_server.go/storage"
)

//Ruan
//Lihle himself #greenpaper
//Ruan adding another comment

func main() {
	err := storage.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := mux.NewRouter()

	// Call functions from handlers
	r.HandleFunc("/", handler.Public)

	// Public
	r.HandleFunc("/devices", handler.ViewPublicDevices)
	r.HandleFunc("/library", handler.ViewPublicLibrary)
	r.HandleFunc("/attendances", handler.ViewPublicAttendances)

	//Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	port := ":3001"
	fmt.Println("http://localhost" + port)
	fmt.Println()
	http.ListenAndServe(port, r)

}

//Lihle Mpapela 212245194
