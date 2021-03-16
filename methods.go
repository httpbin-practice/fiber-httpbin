package main

import (
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetMethod(c *fiber.Ctx) error {
	resp := BaseResp{}
	resp.Origin = c.IP()
	resp.URL = c.OriginalURL()
	args := c.Context().QueryArgs().String()
	headers := c.Request().Header.String()
	querys, _ := url.ParseQuery(args)
	resp.Args = map[string]interface{}{}
	for k, v := range querys {
		resp.Args[k] = v
	}

	headersSlice := strings.Split(headers, "\n")
	resp.Headers = map[string]string{}
	for _, header := range headersSlice {
		headerPair := strings.Split(header, ":")
		if len(headerPair) != 2 {
			continue
		}
		resp.Headers[headerPair[0]] = strings.Trim(headerPair[1], " ")
	}
	return c.JSON(resp)
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
