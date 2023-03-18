package routes

import (
	"fmt"
	"gotemplate/env"

	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func Run() {
	GetRoutes()
	Router.Run(fmt.Sprintf("%s:%s", env.AppUrl, env.Port))
}

func GetRoutes() {
	addPrivateRoutes(Router)
}
