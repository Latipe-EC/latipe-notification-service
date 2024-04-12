package middleware

import (
	"github.com/gofiber/fiber/v2"
	"latipe-notification-service/internal/adapter/authserv"
	"latipe-notification-service/internal/adapter/authserv/dto"
	"latipe-notification-service/pkgUtils/util/errorUtils"
	"strings"
)

type AuthMiddleware struct {
	authServ *authserv.AuthService
}

func NewAuthMiddleware(service *authserv.AuthService) *AuthMiddleware {
	return &AuthMiddleware{authServ: service}
}

func (auth AuthMiddleware) RequiredRoles(roles []string, option ...int) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bearToken := ctx.Get("Authorization")
		if bearToken == "" || len(strings.Split(bearToken, " ")) < 2 {
			return errorUtils.ErrUnauthenticated
		}

		str := strings.Split(bearToken, " ")
		if len(str) < 2 {
			return errorUtils.ErrUnauthenticated
		}

		bearToken = str[1]
		req := dto.AuthorizationRequest{}
		req.Token = bearToken

		resp, err := auth.authServ.Authorization(ctx.Context(), &req)
		if err != nil {
			return errorUtils.ErrInternalServer
		}

		for _, i := range roles {
			if i == resp.Role {
				return ctx.Next()
			}
		}
		return errorUtils.ErrPermissionDenied
	}
}

func (auth AuthMiddleware) RequiredAuthentication() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bearToken := ctx.Get("Authorization")
		if bearToken == "" {
			return errorUtils.ErrUnauthenticated
		}

		str := strings.Split(bearToken, " ")
		if len(str) < 2 {
			return errorUtils.ErrUnauthenticated
		}

		bearToken = str[1]
		req := dto.AuthorizationRequest{
			Token: bearToken,
		}
		resp, err := auth.authServ.Authorization(ctx.Context(), &req)
		if err != nil {
			return errorUtils.ErrInternalServer
		}

		ctx.Locals("user_name", resp.Email)
		ctx.Locals(USER_ID, resp.Id)
		ctx.Locals("bearer_token", bearToken)
		return ctx.Next()
	}
}

func (auth AuthMiddleware) RequiredAPIKeyHeader() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		apiKey := ctx.Get("x-api-key")
		if apiKey == "" {
			return errorUtils.ErrUnauthenticated
		}

		return ctx.Next()
	}
}
