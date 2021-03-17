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
	headerMap := map[string]string{}
	c.Response().Header.VisitAll(func(k []byte, v []byte) {
		headerMap[string(k)] = string(v)
	})
	return c.JSON(headerMap)
}

func PostRespHeaders(c *fiber.Ctx) error {
	headerMap := map[string]string{}
	c.Response().Header.VisitAll(func(k []byte, v []byte) {
		headerMap[string(k)] = string(v)
	})
	return c.JSON(headerMap)
}
