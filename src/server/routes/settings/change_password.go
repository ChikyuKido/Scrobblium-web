package settings

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrobblium-web/database/repo"
	"scrobblium-web/util"
)

func ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			CurrentPassword string `json:"currentPassword" binding:"required"`
			NewPassword     string `json:"newPassword" binding:"required"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}
		currentPasswordHash, _ := repo.GetPassword()
		if !util.CheckPasswordHash(currentPasswordHash, input.CurrentPassword) {
			c.JSON(401, gin.H{
				"error": "Bad Credentials",
			})
			return
		}
		newHashedPassword, _ := util.HashPassword(input.NewPassword)
		repo.UpdatePassword(newHashedPassword)
		c.JSON(200, gin.H{"message": "Changed Password Successfully"})
	}
}
