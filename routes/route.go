package routes

import (
	"golang-chap49/infra"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoutes(r, ctx)
	return r
}

func authRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	authGroup := r.Group("/auth")
	authGroup.POST("/login", ctx.Ctl.User.LoginController)
	authGroup.POST("/check-email", ctx.Ctl.User.CheckEmailUserController)
	authGroup.POST("/register", ctx.Ctl.User.CreateUserController)
	authGroup.PATCH("/reset-password", ctx.Ctl.User.ResetUserPasswordController)
	authGroup.POST("/otp-verify/:otp", ctx.Ctl.User.CheckOTP)
}
