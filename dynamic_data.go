package main

import (
	"encoding/base64"

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

func GetUUID(c *fiber.Ctx) error {
	u, _ := uuid.NewUUID()
	return c.JSON(map[string]string{"uuid": u.String()})
}
