package main

import "github.com/gofiber/fiber/v2"

func GetImage(c *fiber.Ctx) error {
	return c.SendString("get image")
}

func GetImageJPEG(c *fiber.Ctx) error {
	return c.SendString("get image jpeg")
}

func GetImagePNG(c *fiber.Ctx) error {
	return c.SendString("get image png")
}

func GetImageSVG(c *fiber.Ctx) error {
	return c.SendString("get image svg")
}

func GetImageWebP(c *fiber.Ctx) error {
	return c.SendString("get image webp")
}
