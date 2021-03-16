package main

import "github.com/gofiber/fiber/v2"

func DeleteAnything(c *fiber.Ctx) error {
	return c.SendString("delete anything!")
}

func GetAnything(c *fiber.Ctx) error {
	return c.SendString("get anything")
}

func PatchAnything(c *fiber.Ctx) error {
	return c.SendString("Patch anything")
}

func PostAnything(c *fiber.Ctx) error {
	return c.SendString("post anything")
}

func PutAnything(c *fiber.Ctx) error {
	return c.SendString("put anything")
}

func DeleteAnythingAnything(c *fiber.Ctx) error {
	return c.SendString("delete anythign anything")
}

func GetAnythingAnything(c *fiber.Ctx) error {
	return c.SendString("get anything anything")
}

func PatchAnythingAnything(c *fiber.Ctx) error {
	return c.SendString("patch anythign anything")
}

func PostAnythingAnything(c *fiber.Ctx) error {
	return c.SendString("post anything anything")
}

func PutAnythingAnything(c *fiber.Ctx) error {
	return c.SendString("put anything anything")
}
