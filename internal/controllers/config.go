package lftcontrollers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/sedyukov/volunteer-server/internal/storage"
)

func GetexternalConfig(c *fiber.Ctx) error {
	config := storage.GetExternalConfig()
	c.JSON(config)
	return nil
}

func GetOwnConfig(c *fiber.Ctx) error {
	config := storage.GetOwnConfig()
	c.JSON(config)
	return nil
}
