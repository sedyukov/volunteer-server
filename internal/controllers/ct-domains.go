package lftcontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sedyukov/volunteer-server/internal/storage"
)

func GetCtDomains(c *fiber.Ctx) error {
	domains := storage.GetCtDomains()
	c.JSON(domains)
	return nil
}
