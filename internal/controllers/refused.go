package lftcontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sedyukov/volunteer-server/internal/storage"
)

func GetRefused(c *fiber.Ctx) error {
	domains := storage.GetRefused()
	c.JSON(domains)
	return nil
}
