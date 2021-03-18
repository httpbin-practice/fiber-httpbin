package main

import (
	"github.com/gofiber/fiber/v2"
)

func GetHeaders(c *fiber.Ctx) error {
	headers := map[string]string{}
	c.Request().Header.VisitAll(func(k, v []byte) {
		headers[string(k)] = string(v)
	})
	return c.JSON(map[string]interface{}{"headers": headers})
}

func GetIP(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"origin": c.IP()})
}

func GetUserAgent(c *fiber.Ctx) error {
	userAgent := string(c.Context().UserAgent())
	return c.JSON(map[string]string{"User-Agent": userAgent})
}
