package main

import "github.com/gofiber/fiber/v2"

func DeleteStatus(c *fiber.Ctx) error {
	return c.SendString(c.Params("codes"))
}

func GetStatus(c *fiber.Ctx) error {
	return c.SendString(c.Params("codes"))

}

func PatchStatus(c *fiber.Ctx) error {
	return c.SendString(c.Params("codes"))

}

func PostStatus(c *fiber.Ctx) error {
	return c.SendString(c.Params("codes"))

}

func PutStatus(c *fiber.Ctx) error {
	return c.SendString(c.Params("codes"))

}
