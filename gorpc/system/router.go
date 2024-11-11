package system

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()
	handler := NewHandler(db)
	r.GET("/", handler.GetLoaders)
	return r
}
