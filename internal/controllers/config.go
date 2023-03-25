package lftcontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sedyukov/volunteer-server/internal/storage"
)

func GetCentralizedConfig(c *fiber.Ctx) error {
	config := storage.GetCentralizedConfig()
	c.JSON(config)
	return nil
}
