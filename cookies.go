package main

import "github.com/gofiber/fiber/v2"

func GetCookies(c *fiber.Ctx) error {
	return c.SendString("get cookies")
}

func GetCookiesDelete(c *fiber.Ctx) error {
	return c.SendString("get cookies delete")
}

func GetCookiesSet(c *fiber.Ctx) error {
	return c.SendString("get cookies set")
}

func GetCookiesSetValue(c *fiber.Ctx) error {
	return c.SendString("get cookies set value")
}
