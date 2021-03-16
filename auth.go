package main

import "github.com/gofiber/fiber/v2"

func GetBasicAuth(c *fiber.Ctx) error {
	return c.SendString("get basic auth")
}

func GetBearerAuth(c *fiber.Ctx) error {
	return c.SendString("get bearer auth")
}


func GetDigestAuth(c *fiber.Ctx) error {
	return c.SendString("get digest auth")
}


func GetDigestAuthA(c *fiber.Ctx)error{
	return c.SendString("get digest auth a")
}

func GetDigestAuthAS(c *fiber.Ctx)error{
	return c.SendString("get digest auth a s")
}

func GetHidenBasicAuth(c *fiber.Ctx)error{
	return c.SendString("get hiden basic auth")
}