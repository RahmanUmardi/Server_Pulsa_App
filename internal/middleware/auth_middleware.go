package middleware

import (
	"log"
	"net/http"
	"server-pulsa-app/internal/shared/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	RequireToken(roles ...string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtService service.JwtService
}

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func (a *authMiddleware) RequireToken(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var authHeader AuthHeader
		if err := ctx.ShouldBindHeader(&authHeader); err != nil {
			log.Printf("RequireToken: Error binding header: %v \n", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenHeader := strings.TrimPrefix(authHeader.AuthorizationHeader, "Bearer ")
		if tokenHeader == "" {
			log.Println("RequireToken: Missing token")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := a.jwtService.ValidateToken(tokenHeader)
		if err != nil {
			log.Printf("RequireToken: Error parsing token: %v \n", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("employee", claims.UserId)

		role := claims.Role
		if role == "" {
			log.Println("RequireToken: Missing role in token")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !isValidRole(role, roles) {
			log.Println("RequireToken: Invalid role")
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}

func isValidRole(userRole string, validRoles []string) bool {
	for _, role := range validRoles {
		if userRole == role {
			return true
		}
	}
	return false
}

func NewAuthMiddleware(jwtService service.JwtService) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}
