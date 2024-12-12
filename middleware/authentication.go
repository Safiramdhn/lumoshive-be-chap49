package middleware

import (
	"golang-chap49/database"
	"golang-chap49/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Middleware struct {
	log    *zap.Logger
	Cacher database.Cacher
}

func NewMiddleware(log *zap.Logger, cacher database.Cacher) Middleware {
	return Middleware{
		log:    log,
		Cacher: cacher,
	}
}

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			helper.ResponseError(c, "Token is required", "Unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		userID := c.GetHeader("User-ID")
		if userID == "" {
			helper.ResponseError(c, "User-ID is required", "Unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		m.log.Info("Authenticating user", zap.String("userID", userID), zap.String("token", token))

		storedToken, err := m.Cacher.Get(userID)
		if err != nil {
			helper.ResponseError(c, "Failed to retrieve token", "Server error", http.StatusInternalServerError)
			c.Abort()
			return
		}
		m.log.Info("Authenticating user", zap.String("storedToken", storedToken), zap.String("token", token))

		if storedToken == "" || storedToken != token {
			helper.ResponseError(c, "Invalid token", "Unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}

func (m *Middleware) RoleAuthorization(requiredRoles string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("User-Role")
		if role == "" {
			helper.ResponseError(c, "Roles are required", "Unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		m.log.Info("Authorizing user role", zap.String("role", role), zap.String("requiredRoles", requiredRoles))

		if role == requiredRoles {
			c.Next()
			return
		}

		helper.ResponseError(c, "Insufficient permissions", "Unauthorized", http.StatusUnauthorized)
		c.Abort()
	}
}
