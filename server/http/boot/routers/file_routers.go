package routers

import (
	"github.com/ecommerce-service/product-service/server/http/handlers"
	"github.com/gofiber/fiber/v2"
)

type FileRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.HandlerContract
}

func NewFileRouters(routeGroup fiber.Router, handler handlers.HandlerContract) IRouters {
	return &FileRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (r FileRouters) RegisterRouter() {
	handler := handlers.NewFileHandler(r.Handler)

	fileRouters := r.RouteGroup.Group("/file")
	fileRouters.Get("/:key", handler.GetUrlByKey)
	fileRouters.Post("", handler.Upload)
}
