package library

import (
	"github.com/gofiber/fiber/v2"
)

func GetOwnerIdAndTeamId(c *fiber.Ctx) (int, int, error) {

	OwnerId := 1000
	TeamId := 1000

	return OwnerId, TeamId, nil
}
