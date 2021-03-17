package main

import (
	"github.com/gofiber/fiber/v2"
)

func GetHeaders(c *fiber.Ctx) error {
	return c.SendString("get headers")
}

func GetIP(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"origin": c.IP()})
}

func GetUserAgent(c *fiber.Ctx) error {
	userAgent := string(c.Context().UserAgent())
	return c.JSON(map[string]string{"User-Agent": userAgent})
}
