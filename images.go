package main

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetImage(c *fiber.Ctx) error {
	p := "static/img.jpeg"
	accept := c.Accepts("image/webp", "image/png", "image/jpeg", "image/svg")
	if accept != "" {
		fmt.Println(accept)
		p = fmt.Sprintf("static/img.%s", strings.Split(accept, "/")[1])
	}
	return c.SendFile(p, true)
}

func GetImageJPEG(c *fiber.Ctx) error {
	return c.SendFile("static/img.jpeg", true)
}

func GetImagePNG(c *fiber.Ctx) error {
	return c.SendFile("static/img.png", true)
}

func GetImageSVG(c *fiber.Ctx) error {
	return c.SendFile("static/img.svg", true)
}

func GetImageWebP(c *fiber.Ctx) error {
	return c.SendFile("static/img.webp", true)
}
