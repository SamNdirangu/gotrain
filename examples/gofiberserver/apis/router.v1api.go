package apis

import (
	"training/examples/gofiberserver/apis/books"
	"training/examples/gofiberserver/apis/products"
	"training/examples/gofiberserver/apis/users"

	"github.com/gofiber/fiber/v2"
)

func RegisterV1APIs(api fiber.Router) {
	products.RegisterProductRoutes(api)
	books.RegisterBookRoutes(api)
	users.RegisterUserRoutes(api)
}
