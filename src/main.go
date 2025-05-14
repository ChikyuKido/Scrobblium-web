package main

import (
	"scrobblium-web/database"
	"scrobblium-web/server"
)

func main() {
	database.Init()
	server.Init()
}
