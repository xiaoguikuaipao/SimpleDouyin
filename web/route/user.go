package route

import (
	"SimpleDouyin/web/api"

	"github.com/cloudwego/hertz/pkg/route"
)

func registerUserRoute(g *route.RouterGroup) {
	user := g.Group("/user")

	user.POST("/register", api.RegisterHandler)
	user.POST("/login", api.LoginHandler)
}
