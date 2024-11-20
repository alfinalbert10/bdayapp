package handlers

import (
	"bdayappbknd/common"
	"bdayappbknd/managers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	groupName   string
	userManager *managers.UserManager
}

func NewUserHandlerFrom(userManager *managers.UserManager) *UserHandler {
	return &UserHandler{
		"api/users",
		userManager,
	}
}

func (userHandler *UserHandler) RegisterUserApis(r *gin.Engine) {
	userGroup := r.Group(userHandler.groupName)
	userGroup.POST("", userHandler.Create)
	userGroup.GET("", userHandler.List)
	userGroup.GET(":userid/", userHandler.Detail)
	userGroup.DELETE(":userid/", userHandler.Delete)
	userGroup.PATCH(":userid/", userHandler.Update)
}

func (userHandler *UserHandler) Create(ctx *gin.Context) {

	userData := common.NewUserCreationInput()
	err := ctx.BindJSON(&userData)
	if err != nil {
		fmt.Println("failed to bind data")
	}
	newUser, err := userHandler.userManager.Create(userData)

	if err != nil {
		fmt.Println("failed to create user")
	}
	ctx.JSON(200, newUser)
}

func (userHandler *UserHandler) List(ctx *gin.Context) {

	allUser, err := userHandler.userManager.List()

	if err != nil {
		fmt.Println("failed to list users")
	}
	ctx.JSON(200, allUser)
}

func (userHandler *UserHandler) Delete(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("userId")

	if !ok {
		fmt.Println("invalid userId")
	}
	err := userHandler.userManager.Delete(userId)

	if err != nil {
		fmt.Println("failed to delete user")
	}
	common.SucessResponse(ctx, "deleted the user successfully")
}

func (userHandler *UserHandler) Detail(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("userId")

	if !ok {
		fmt.Println("invalid userId")
	}
	user, err := userHandler.userManager.Get(userId)

	if err != nil {
		fmt.Println("failed to get user")
	}
	ctx.JSON(http.StatusOK, user)
}

func (userHandler *UserHandler) Update(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("userId")

	if !ok {
		fmt.Println("invalid userId")
	}

	userData := common.NewUserUpdateInput()

	err := ctx.BindJSON(&userData)

	if err != nil {
		common.BadResponse(ctx, "failed to bind data")
		return
	}
	user, err := userHandler.userManager.Update(userId, userData)

	if err != nil {
		common.BadResponse(ctx, "failed to update user")
		return
	}
	ctx.JSON(http.StatusOK, user)
}
