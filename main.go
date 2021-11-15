package main

import (
	"github.com/douglasandradeee/web_api_gin/database"

	"github.com/douglasandradeee/web_api_gin/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
