package lftcontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sedyukov/volunteer-server/internal/storage"
)

func GetBlockedDomains(c *fiber.Ctx) error {
	domains := storage.GetDomains()
	c.JSON(domains)
	return nil
}
