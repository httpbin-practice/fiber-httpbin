package main

import "github.com/gofiber/fiber/v2"

func GetHeaders(c *fiber.Ctx) error {
	return c.SendString("get headers")
}

func GetIP(c *fiber.Ctx) error {
	return c.SendString("get ip")
}

func GetUserAgent(c *fiber.Ctx) error {
	return c.SendString("get user agent")
}
