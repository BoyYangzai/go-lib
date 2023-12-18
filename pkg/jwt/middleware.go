package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头或其他地方获取令牌
		authorizationHeader := c.GetHeader("Authorization")

		// 检查 Authorization 头是否存在
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 检查 Authorization 头的格式是否正确
		if len(authorizationHeader) < len("Bearer ") || authorizationHeader[:len("Bearer ")] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization format"})
			c.Abort()
			return
		}

		// 获取实际的 JWT 令牌
		token := authorizationHeader[len("Bearer "):]

		// 进行认证检查
		user, err := ParseJwt(token)
		if err != nil {
			// 如果令牌无效，返回未认证的错误
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 如果令牌有效，继续处理请求
		c.Set("user", user)
		println(user.Username + " is authorized\n")
		c.Next()
	}
}

func LoginAuthStorage(c *gin.Context) (bool, string) {
	// 从请求头或其他地方获取令牌
	authorizationHeader := c.GetHeader("Authorization")

	// 检查 Authorization 头是否存在
	if authorizationHeader == "" {
		c.Abort()
		return false, ""
	}

	// 检查 Authorization 头的格式是否正确
	if len(authorizationHeader) < len("Bearer ") || authorizationHeader[:len("Bearer ")] != "Bearer " {
		c.Abort()
		return false, ""
	}

	// 获取实际的 JWT 令牌
	token := authorizationHeader[len("Bearer "):]

	// 进行认证检查
	user, err := ParseJwt(token)
	if err != nil {
		// 如果令牌无效，返回未认证的错误
		return false, ""
	}

	// 如果令牌有效，继续处理请求
	c.Set("user", user)
	println(user.Username + " is authorized\n")
	return true, token
}
