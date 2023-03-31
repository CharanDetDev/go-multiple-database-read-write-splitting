package caller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func Response(c *fiber.Ctx, httpCode int, data interface{}) (err error) {

	js, _ := json.Marshal(data)
	c.Write(js)
	c.Status(httpCode)

	return nil
}
