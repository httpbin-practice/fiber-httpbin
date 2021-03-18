package main

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func GetRespContent(c *fiber.Ctx) BaseResp {
	resp := BaseResp{}
	resp.Method = c.Method()
	resp.Origin = c.IP()
	resp.URL = c.Context().URI().String()
	resp.Headers = map[string]string{}
	c.Request().Header.VisitAll(func(k, v []byte) {
		resp.Headers[string(k)] = string(v)
	})
	args := c.Context().QueryArgs().String()
	querys, _ := url.ParseQuery(args)
	resp.Args = map[string]interface{}{}
	for k, v := range querys {
		resp.Args[k] = v
	}
	return resp
}

func GetMethod(c *fiber.Ctx) error {
	return c.JSON(GetRespContent(c))
}

func PostMethod(c *fiber.Ctx) error {
	return c.JSON(GetRespContent(c))
}

func PutMethod(c *fiber.Ctx) error {
	return c.JSON(GetRespContent(c))
}

func PatchMethod(c *fiber.Ctx) error {
	return c.JSON(GetRespContent(c))
}

func DeleteMethod(c *fiber.Ctx) error {
	return c.JSON(GetRespContent(c))
}
