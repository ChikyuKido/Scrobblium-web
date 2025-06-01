package auth

import (
	"github.com/gin-gonic/gin"
	"scrobblium-web/database/repo"
	"scrobblium-web/util"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginResponse = struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}
		err := c.BindJSON(&loginResponse)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "invalid login request",
			})
			return
		}
		username, _ := repo.GetUsername()
		if username != loginResponse.Username {
			c.JSON(401, gin.H{
				"error": "Bad Credentials",
			})
			return
		}
		password, _ := repo.GetPassword()
		if !util.CheckPasswordHash(password, loginResponse.Password) {
			c.JSON(401, gin.H{
				"error": "Bad Credentials",
			})
			return
		}

		jwt, err := util.GenerateJWT(username)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "internal server error",
			})
			return
		}
		c.JSON(200, gin.H{
			"jwt": jwt,
		})
	}
}
