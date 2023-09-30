package route

import (
	"order/config"
	"order/controller"
	"order/middleware"
	"order/repository"
	"order/service"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	db := config.InitDatabase()

	api := app.Group("/api")

	orderRepository := repository.NewOrderRepositoryPG(db)
	orderService := service.NewOrderService(orderRepository)
	orderController := controller.NewOrderController(orderService)

	api.Get("/order/inventory", middleware.Protected(), orderController.GetListAllInventory)
	api.Get("/order/:id/inventory", middleware.Protected(), orderController.GetListInventory)
	api.Post("/order/inventory", middleware.Protected(), orderController.CreateInventory)
	api.Put("/order/:id/inventory", middleware.Protected(), orderController.UpdateInventory)
	api.Delete("/order/:id/inventory", middleware.Protected(), orderController.DeleteInventory)

}
