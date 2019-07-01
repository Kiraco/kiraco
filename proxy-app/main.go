package main

import (
	handlers "github.com/kiraco/proxy-app/api/handlers"
	server "github.com/kiraco/proxy-app/api/server"
	utils "github.com/kiraco/proxy-app/api/utils"
)

func main() {
	/*
		Router = Iris
		Env vars
	*/
	utils.LoadEnv()
	app := server.SetUp()
	handlers.HandleRedirection(app)
	server.RunServer(app)

}
