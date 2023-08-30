package main

import (
	"SimpleDouyin/web/route"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
)

func main() {
	h := server.Default(
		server.WithHostPorts("127.0.0.1:8080"),
		server.WithTransport(standard.NewTransporter),
	)

	//Initialize the route
	route.RegisterRoute(h)

	h.Spin()
}
