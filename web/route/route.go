package route

import "github.com/cloudwego/hertz/pkg/app/server"

func RegisterRoute(h *server.Hertz) {
	douyin := h.Group("/douyin")

	registerUserRoute(douyin)

	registerCommentRoute(douyin)

	registerFavourRoute(douyin)

	registerSocialRoute(douyin)

	registerVideoRoute(douyin)
}
