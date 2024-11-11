package system

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Controller *Controller
}

func NewHandler(db *sql.DB) *Handler {
	controller := NewController(db)
	return &Handler{Controller: controller}
}

func (h *Handler) GetLoaders(c *gin.Context) {
	h.Controller.GetLoaders(c)
}
