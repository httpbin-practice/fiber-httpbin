package main

import (
	"encoding/base64"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetBase64Value decode base64 value
func GetBase64Value(c *fiber.Ctx) error {
	respText := "Incorrect Base64 data try: SFRUUEJJTiBpcyBhd2Vzb21l"
	result, err := base64.StdEncoding.DecodeString(c.Params("value"))
	if err == nil {
		respText = string(result)
	}
	return c.SendString(respText)
}

// GetBytesN x
func GetBytesN(c *fiber.Ctx) error {
	// n, err := strconv.Atoi(c.Params("n"))
	// if err != nil {
	// 	return err
	// }
	return c.SendString("send get bytes N")
}

func Delay(c *fiber.Ctx) error {
	n, err := strconv.Atoi(c.Params("delay"))
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(n) * time.Second)
	return c.JSON(GetRespContent(c))
}

func GetDrip(c *fiber.Ctx) error {
	if c.Request().Header.Peek("If-Modified-Since") != nil || c.Request().Header.Peek("If-None-Match") != nil {
		resp := GetRespContent(c)
		c.Response().Header.Set("Last-Modified", "http-date")
		c.Response().Header.Set("ETag", "uuid")
		return c.JSON(resp)
	}
	c.Response().SetStatusCode(304)
	return c.Send(nil)
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

func GetUUID(c *fiber.Ctx) error {
	u, _ := uuid.NewUUID()
	return c.JSON(map[string]string{"uuid": u.String()})
}
