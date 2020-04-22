package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sundowndev/phoneinfoga/pkg/scanners"
)

// JSONResponse is the default API response type
type JSONResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type scanURL struct {
	Number uint `uri:"number" binding:"required,min=2"`
}

// ValidateScanURL validates scan URLs
func ValidateScanURL(c *gin.Context) {
	var v scanURL

	if err := c.ShouldBindUri(&v); err != nil {
		errorHandling(c, "Parameter 'number' must be a valid integer.")
		return
	}

	number, err := scanners.LocalScan(c.Param("number"))

	if err != nil {
		errorHandling(c, err.Error())
		return
	}

	c.Set("number", number)
}

func errorHandling(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, JSONResponse{Success: false, Error: msg})
	c.Abort()
}
