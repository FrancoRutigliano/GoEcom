package main

import (
	server "GoEcom/infrastructure/entrypoint/server_http"
	"log"
)

func init() {

}

func main() {
	s := server.NewServer()

	if err := s.Run(); err != nil {
		log.Fatal("error running server: ", err)
	}
}
