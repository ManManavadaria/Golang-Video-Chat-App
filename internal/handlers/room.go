package handlers

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

func RoomCreate(c *fiber.Ctx) error {
	c.Redirect(fmt.Sprintf("/room/%s", uuid.New().String()))
	return nil
}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(404)
		return nil
	}
}

func RoomWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")

	if uuid == "" {
		return
	}
}

func CreateOrGetRoom(uuid string) (string, string, Room) {

}
