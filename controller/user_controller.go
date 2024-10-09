package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reynaldineo/go-gin-gorm-starter/constant"
	"github.com/reynaldineo/go-gin-gorm-starter/dto"
	"github.com/reynaldineo/go-gin-gorm-starter/service"
	"github.com/reynaldineo/go-gin-gorm-starter/utils"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		GetMe(ctx *gin.Context)
	}

	userController struct {
		userService service.UserService
	}
)

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var user dto.UserRegisterRequest
	if err := ctx.ShouldBind(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userRes, err := c.userService.RegisterUser(ctx.Request.Context(), user)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, userRes)
	ctx.JSON(http.StatusCreated, resp)
}

func (c *userController) Login(ctx *gin.Context) {
	var req dto.UserLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.VerifyUser(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	resp := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN, result)
	ctx.JSON(http.StatusOK, resp)
}

func (c *userController) GetMe(ctx *gin.Context) {
	userId := ctx.MustGet(constant.CTX_KEY_USER_ID).(string)

	user, err := c.userService.GetUserByID(ctx.Request.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, user)
	ctx.JSON(http.StatusOK, resp)
}
