package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListSnackHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GETALL to UaiFOOD",
	})
}
