package main

import "github.com/gofiber/fiber/v2"

func GetRespBrotli(c *fiber.Ctx) error {
	return c.SendString("get brotli response")
}

func GetRespDeflate(c *fiber.Ctx) error {
	return c.SendString("get Deflate response")
}

func GetRespDeny(c *fiber.Ctx) error {
	return c.SendString("get deny response")
}

func GetRespEncodingUTF8(c *fiber.Ctx) error {
	return c.SendString("get encoding utf8 response")
}

func GetRespGZip(c *fiber.Ctx) error {
	return c.SendString("get gzip response")
}

func GetRespHTML(c *fiber.Ctx) error {
	return c.SendString("get html response")
}

func GetRespJSON(c *fiber.Ctx) error {
	return c.SendString("get JSON response")
}

func GetRespRobotsTXT(c *fiber.Ctx) error {
	return c.SendString("get robots.txt response")
}

func GetRespXML(c *fiber.Ctx) error {
	return c.SendString("get xml response")
}
