package main

import "github.com/gofiber/fiber/v2"

func GetAbsoluteRedirect(c *fiber.Ctx) error {
	return c.SendString("get absolute redirect")
}

func DeleteRedirectTo(c *fiber.Ctx) error {
	return c.SendString("delete redirect to")
}

func GetRedirectTo(c *fiber.Ctx) error {
	return c.SendString("get redirect to")
}

func PatchRedirectTo(c *fiber.Ctx) error {
	return c.SendString("patch redirect to")
}

func PostRedirectTo(c *fiber.Ctx) error {
	return c.SendString("post redirect to")
}

func PutRedirectTo(c *fiber.Ctx) error {
	return c.SendString("put redirect to")
}

func GetRedirect(c *fiber.Ctx) error {
	return c.SendString("get redirect")
}

func GetRelativeRedirect(c *fiber.Ctx) error {
	return c.SendString("get relative redirect")
}
