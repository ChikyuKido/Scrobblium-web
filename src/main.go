package main

import (
	"scrobblium-web/database"
	"scrobblium-web/server"
	"scrobblium-web/util"
)

func main() {
	database.Init()
	util.AddDefaultSettings()
	server.Init()
}
