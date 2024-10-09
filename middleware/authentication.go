package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/reynaldineo/go-gin-gorm-starter/constant"
	"github.com/reynaldineo/go-gin-gorm-starter/dto"
	"github.com/reynaldineo/go-gin-gorm-starter/service"
	"github.com/reynaldineo/go-gin-gorm-starter/utils"
)

func Authenticate(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_VERIFY_TOKEN, dto.ErrTokenNotFound.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !strings.Contains(authHeader, "Bearer ") {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_VERIFY_TOKEN, dto.ErrTokenInvalid.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		userId, userRole, err := jwtService.GetPayloadInsideToken(authHeader)
		if err != nil {
			if err.Error() == dto.ErrTokenExpired.Error() {
				respose := utils.BuildResponseFailed(dto.MESSAGE_FAILED_VERIFY_TOKEN, dto.ErrTokenExpired.Error(), nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, respose)
				return
			}
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_VERIFY_TOKEN, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Set(constant.CTX_KEY_ROLE_NAME, authHeader)
		ctx.Set(constant.CTX_KEY_USER_ID, userId)
		ctx.Set(constant.CTX_KEY_ROLE_NAME, userRole)
		ctx.Next()
	}
}
