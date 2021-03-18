package main

import (
	"github.com/gofiber/fiber/v2"
)

func GetCookies(c *fiber.Ctx) error {
	cookies := map[string]string{}
	c.Request().Header.VisitAllCookie(func(k, v []byte) {
		cookies[string(k)] = string(v)
	})
	return c.JSON(map[string]interface{}{"cookies": cookies})
}

func GetCookiesDelete(c *fiber.Ctx) error {
	return c.SendString("get cookies delete")
}

func GetCookiesSet(c *fiber.Ctx) error {
	cookies := map[string]string{}
	c.Request().Header.VisitAllCookie(func(k, v []byte) {
		cookies[string(k)] = string(v)
	})
	c.Context().QueryArgs().VisitAll(func(k, v []byte) {
		name, value := string(k), string(v)
		c.Cookie(&fiber.Cookie{
			Name:  name,
			Value: value,
		})
		cookies[name] = value
	})

	return c.JSON(map[string]interface{}{"cookies": cookies})
}

func GetCookiesSetValue(c *fiber.Ctx) error {
	name, value := c.Params("name"), c.Params("value")
	cookies := map[string]string{}
	c.Cookie(&fiber.Cookie{
		Name:  name,
		Value: value,
	})
	c.Request().Header.VisitAllCookie(func(k, v []byte) {
		cookies[string(k)] = string(v)
	})
	cookies[name] = value
	return c.JSON(map[string]interface{}{"cookies": cookies})
}
