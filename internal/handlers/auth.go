package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	openapi "github.com/rodpk/library-mgmt-api/api"
)

type AuthHandler struct {

}

func Register(c *gin.Context) {
	var user openapi.UserRegisterRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	
}