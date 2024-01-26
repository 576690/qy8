package main

import (
	"log"
	"net/http"

	"github.com/576690/qy8/backend/router"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	router.Router()

	error := server.ListenAndServe()
	if error != nil {
		log.Println(error)
	}
	// 启动服务器

}
