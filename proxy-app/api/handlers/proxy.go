package handlers

import (
	"encoding/json"

	"github.com/kataras/iris"
	"github.com/kiraco/proxy-app/api/middleware"
)

// HandleRedirection - Handle routing redirection
func HandleRedirection(app *iris.Application) {
	app.Get("/ping", middleware.ProxyMiddleware, proxyHandler)
}

func proxyHandler(c iris.Context) {
	response, err := json.Marshal(middleware.QueueList)
	if err != nil {
		c.JSON(iris.Map{"status": 400, "result": "parse error"})
		return
	}
	c.JSON(iris.Map{"result": string(response)})

}
