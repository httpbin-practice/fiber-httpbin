package main

import "github.com/gofiber/fiber/v2"

func DeleteStatus(c *fiber.Ctx) error {
	return c.SendString("delete status")
}

func GetStatus(c *fiber.Ctx) error {
	return c.SendString("get status")
}

func PatchStatus(c *fiber.Ctx) error {
	return c.SendString("patch status")
}

func PostStatus(c *fiber.Ctx) error {
	return c.SendString("post status")
}

func PutStatus(c *fiber.Ctx) error {
	return c.SendString("put status")
}
