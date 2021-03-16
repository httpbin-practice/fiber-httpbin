package main

import "github.com/gofiber/fiber/v2"

func GetRespCache(c *fiber.Ctx) error {
	return c.SendString("get resp cache")
}

func GetRespCacheValue(c *fiber.Ctx) error {
	return c.SendString("get resp cache value")
}

func GetRespEtag(c *fiber.Ctx) error {
	return c.SendString("get resp etag")
}

func GetRespHeaders(c *fiber.Ctx) error {
	return c.SendString("get resp headers")
}

func PostRespHeaders(c *fiber.Ctx) error {
	return c.SendString("post resp headers")
}
