package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAbsoluteRedirect(c *fiber.Ctx) error {
	return c.SendString("get absolute redirect")
}

func RedirectTo(c *fiber.Ctx) error {
	code, _ := strconv.Atoi(c.Query("code"))
	return c.Redirect(c.Query("url"), code)
}

func GetRelativeRedirect(c *fiber.Ctx) error {
	return c.SendString("get relative redirect")
}
