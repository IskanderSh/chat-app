package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/udborets/chat-app/server/internal/models"
	"github.com/udborets/chat-app/server/internal/responses"
	"github.com/udborets/chat-app/server/internal/service"
	"net/http"
)

type AuthHTTP struct {
	authBLogic service.IAuthBLogic
}

func NewHTTP() *AuthHTTP {
	return &AuthHTTP{
		authBLogic: service.NewAuthBLogic(),
	}
}

func (h *AuthHTTP) InitRoutes(router *gin.Engine) {
	authAPI := router.Group("/auth")

	authAPI.POST("/signup", h.userSignUp)
	authAPI.POST("/signin", h.userSignIn)
	authAPI.GET("/validate", h.requireAuth, h.validate)
}

func (h *AuthHTTP) userSignUp(ctx *gin.Context) {
	var inp models.UserSignUpInput

	if err := ctx.BindJSON(&inp); err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, "invalid input body", err)
		return
	}

	statusCode, msg, err := h.authBLogic.SignUp(inp)
	responses.NewResponse(ctx, statusCode, msg, err)
}

func (h *AuthHTTP) userSignIn(ctx *gin.Context) {
	var inp models.UserSignInInput

	if err := ctx.BindJSON(&inp); err != nil {
		responses.NewResponse(ctx, http.StatusBadRequest, "invalid input body", err)
		return
	}

	statusCode, msg, err := h.authBLogic.SignIn(inp)
	if err != nil {
		responses.NewResponse(ctx, statusCode, msg, err)
		return
	}

	jwtToken := msg
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", jwtToken, 3600*24*30, "", "", false, true)
}

func (h *AuthHTTP) requireAuth(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		responses.NewResponse(ctx, http.StatusUnauthorized, err.Error(), err)
		return
	}

	user, statusCode, msg, err := h.authBLogic.ParseJWTToken(tokenString)
	if err != nil {
		responses.NewResponse(ctx, statusCode, msg, err)
		return
	}

	ctx.Set("user", user.(models.User))

	ctx.Next()
}

func (h *AuthHTTP) validate(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"message": user})
}
