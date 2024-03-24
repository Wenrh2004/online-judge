// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
	"online-judge/biz/handler"
	_ "online-judge/docs"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
	url := swagger.URL("http://localhost:8888/swagger/doc.json")
	r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
}
