package middlewares

import (
	"nls-go-messaging/internal/constants"
	"nls-go-messaging/internal/handlers/database"

	"github.com/gofiber/fiber/v2"
)

func RequestDBMiddleware(c *fiber.Ctx) error {
	//application := c.Context().Value(constants.ApplicationCtx).(constants.AppKey)
	DB := c.Locals(constants.DBLocals).(*database.PG_DB)

	if c != nil {
		c.Locals(constants.RequestDBLocals, DB)
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}
	return c.Next()
}
