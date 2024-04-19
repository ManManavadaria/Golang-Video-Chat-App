package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ManManavadaria/Golang-Video-Chat-App/server"
)

func main() {

	http.HandleFunc("/create", server.CreateRoomHandler)
	http.HandleFunc("/join", server.JoinRoomReqHandler)

	port := ":8000"
	fmt.Println("server is running on port", port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal("server error")
	}
}
