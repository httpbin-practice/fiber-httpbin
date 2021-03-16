package main

import "github.com/gofiber/fiber/v2"

func GetMethod(c *fiber.Ctx) error {
	return c.SendString("get methods!")
}

func PostMethod(c *fiber.Ctx) error {
	return c.SendString("post methods!")
}

func PutMethod(c *fiber.Ctx) error {
	return c.SendString("put methods!")
}

func PatchMethod(c *fiber.Ctx) error {
	return c.SendString("patch methods!")
}

func DeleteMethod(c *fiber.Ctx) error {
	return c.SendString("delete methods!")
}

