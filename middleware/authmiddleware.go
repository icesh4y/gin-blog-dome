package middleware

import (
	"essential/dao"
	"essential/jwt"
	"essential/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusForbidden, gin.H{"code": 402, "msg": "权限不足"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := jwt.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "token无效"})
			c.Abort()
			return
		}

		var user models.User
		dao.DB.First(&user, claims.UserId)
		if user.ID == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "token无效"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}

}
