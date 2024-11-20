package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserCreationInput struct {
	Username  string `json:"username"`
	Birthdate string `json:"birthdate"`
}
type UserUpdateInput struct {
	Username  string `json:"username"`
	Birthdate string `json:"birthdate"`
}

func NewUserUpdateInput() *UserUpdateInput {
	return &UserUpdateInput{}
}

func NewUserCreationInput() *UserCreationInput {
	return &UserCreationInput{}
}

type requestResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func SucessResponse(ctx *gin.Context, msg string) {
	response := requestResponse{
		msg,
		strconv.Itoa(http.StatusOK),
	}
	ctx.JSON(http.StatusOK, response)
}

func BadResponse(ctx *gin.Context, msg string) {
	response := requestResponse{
		msg,
		strconv.Itoa(http.StatusBadRequest),
	}
	ctx.JSON(http.StatusBadRequest, response)
}
