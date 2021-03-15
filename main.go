package main

import "dashboard-api/server/handlers"

func main() {
	server := handlers.HTTPServer{
		ServerConf: handlers.ServerConf{
			Address:       ":1231",
		},
	}

	server.Run()
}