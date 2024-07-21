package main

import serverhttp "GoEcom/shared/infrastructure/entrypoint/server_http"

func init() {

}

func main() {
	server := serverhttp.NewServer()

	err := server.Run()
	if err != nil {
		panic(err)
	}
}
