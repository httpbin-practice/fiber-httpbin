package main

import "github.com/gofiber/fiber/v2"

func GetBase64Value(c *fiber.Ctx) error {
	return c.SendString("send get base64 value")
}

func GetBytesN(c *fiber.Ctx) error {
	return c.SendString("send get bytes N")
}

func GetDelay(c *fiber.Ctx) error {
	return c.SendString("get delay")
}

func DeleteDelay(c *fiber.Ctx) error {
	return c.SendString("delete delay")
}

func PostDelay(c *fiber.Ctx) error {
	return c.SendString("post delay")
}
func PatchDelay(c *fiber.Ctx) error {
	return c.SendString("patch delay")
}
func PutDelay(c *fiber.Ctx) error {
	return c.SendString("put delay")
}

func GetDrip(c *fiber.Ctx) error {
	return c.SendString("get drip")
}

func GetLinks(c *fiber.Ctx) error {
	return c.SendString("get links")
}

func GetRange(c *fiber.Ctx) error {
	return c.SendString("get range")
}

func GetStreamBytes(c *fiber.Ctx) error {
	return c.SendString("get stream bytes")
}

func GetStream(c *fiber.Ctx) error {
	return c.SendString("get stream ")
}

func GetUUID(c *fiber.Ctx)error{
	return c.SendString("get uuid")
}