package middlewares

import (
	"strings"

	"github.com/dmzsz/duozhuayu/internal/constants"
	"github.com/dmzsz/duozhuayu/pkg/jwt"
	"github.com/gin-gonic/gin"

	// "strings"
	// "github.com/dmzsz/duozhuayu/internal/constants"
	V1Handler "github.com/dmzsz/duozhuayu/internal/http/handlers/v1"
	// V1JWTService "github.com/dmzsz/duozhuayu/internal/services/jwtservice/v1"
	// "github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	roles []string
}

func NewAuthMiddleware(roles []string) gin.HandlerFunc {
	return (&AuthMiddleware{
		// jwtService: jwtService,
		roles: roles,
	}).Handle
}

func (m *AuthMiddleware) Handle(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		V1Handler.NewAbortResponse(ctx, "missing authorization header")
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		V1Handler.NewAbortResponse(ctx, "invalid header format")
		return
	}

	if headerParts[0] != "Bearer" {
		V1Handler.NewAbortResponse(ctx, "token must content bearer")
		return
	}

	user, err := jwt.ParseToken(headerParts[1])
	if err != nil {
		V1Handler.NewAbortResponse(ctx, "invalid token")
		return
	}

	if IsContain(user.RoleIds, m.roles) && len(user.RoleIds) == 0 {
		V1Handler.NewAbortResponse(ctx, "you don't have access for this action")
		return
	}

	ctx.Set(constants.CtxAuthenticatedUserKey, user)
	ctx.Next()
}

func IsContain(items []string, items2 []string) bool {
	for _, eachItem := range items {
		for _, item2 := range items2 {
			if eachItem == item2 {
				return true
			}
		}
	}
	return false
}
