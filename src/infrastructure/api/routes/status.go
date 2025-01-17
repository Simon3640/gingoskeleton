package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health Check
// @Summary
// @Schemes
// @Tags Health Check
// @Accept json
// @Produce json
// @Router /api/health-check [get]
func getHealthCheck(c *gin.Context) {
	healthCheck := "Hello World"
	c.JSON(http.StatusOK, healthCheck)
}
