package controller

import (
	"order/library"
	"order/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) orderController {
	return orderController{orderService: orderService}
}

func (h orderController) GetListAllInventory(c *fiber.Ctx) error {
	OwnerId, TeamId, err := library.GetOwnerIdAndTeamId(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	params := service.ParamsFilterInventory{
		Q:       c.Query("q"),
		Limit:   limit,
		Page:    page,
		OwnerId: OwnerId,
		TeamId:  TeamId,
	}

	response, err := h.orderService.GetListAllInventory(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": response,
	})

}

func (h orderController) CreateInventory(c *fiber.Ctx) error {

	OwnerId, TeamId, err := library.GetOwnerIdAndTeamId(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	request := service.InventoryRequest{}

	err = c.BodyParser(&request)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := h.orderService.CreateInventory(request, OwnerId, TeamId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"data": response,
	})

}

func (h orderController) GetListInventory(c *fiber.Ctx) error {

	OwnerId, TeamId, err := library.GetOwnerIdAndTeamId(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	inventoryId, _ := strconv.Atoi(c.Params("id"))
	order, err := h.orderService.GetListInventory(inventoryId, OwnerId, TeamId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": order,
	})
}

func (h orderController) DeleteInventory(c *fiber.Ctx) error {

	OwnerId, TeamId, err := library.GetOwnerIdAndTeamId(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	inventoryId, _ := strconv.Atoi(c.Params("id"))

	response, err := h.orderService.DeleteInventory(inventoryId, OwnerId, TeamId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"data": response,
	})
}

func (h orderController) UpdateInventory(c *fiber.Ctx) error {

	OwnerId, TeamId, err := library.GetOwnerIdAndTeamId(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	id, _ := strconv.Atoi(c.Params("id"))
	request := service.InventoryRequest{}
	err = c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := h.orderService.UpdateInventory(request, id, OwnerId, TeamId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"data": response,
	})

}
