package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "test/api/docs"
	"test/api/handler"
	"test/pkg/logger"
	"test/service"
)

// @title           Auth API
// @version         1.0
// @description     Authentication and Authorization API
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(services service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.New(services, log)
	r := gin.Default()

	// === OTP ===
	r.POST("/otp/send", h.SendOTP)
	r.POST("/otp/confirm", h.ConfirmOTP)

	// === Auth ===
	r.POST("/signup", h.SignUp)
	r.POST("/login", h.Login)

	role := r.Group("/role")
	role.Use(h.AuthorizerMiddleware) // 🔐 faqat token bo‘lsa ishlaydi
	{
		role.POST("/", h.CreateRole)
		role.PUT("/:id", h.UpdateRole)
		role.GET("/", h.ListRoles)
	}

	sysuser := r.Group("/sysuser")
	sysuser.Use(h.AuthorizerMiddleware)
	{
		sysuser.POST("/", h.CreateSysUser)
	}
	// === Swagger ===
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
