package server

import (
	"os"

	"github.com/kataras/iris"
)

// SetUp - sets up the iris server
func SetUp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	return app
}

// RunServer -runs the set up'ed server
func RunServer(app *iris.Application) {
	app.Run(
		iris.Addr(os.Getenv("PORT")),
	)
}
