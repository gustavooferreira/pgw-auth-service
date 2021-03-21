package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *Server) ValidateUser(c *gin.Context) {
	bodyData := struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	err := c.ShouldBindJSON(&bodyData)
	if err != nil {
		s.Logger.Info(fmt.Sprintf("error parsing body: %s", err.Error()))
		RespondWithError(c, 400, "error parsing body")
		return
	}

	valid := s.Repo.ValidateUser(bodyData.Username, bodyData.Password)

	c.JSON(200, gin.H{"valid": valid})
}
