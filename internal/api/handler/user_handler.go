package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"project/internal/service"
)

type UserHandler struct {
	service *service.UserService
	logger  *zap.Logger
}

func NewUserHandler(service *service.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

// GetUser godoc
// @Summary      Get user by ID
// @Description  Returns user info by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  model.User
// @Failure      400  {object}  map[string]string
// @Router       /users/{id} [get]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		h.logger.Warn("invalid user id param",
			zap.String("id", c.Params("id")),
		)
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		h.logger.Warn("get user failed",
			zap.Int("id", id),
			zap.Error(err),
		)
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}
