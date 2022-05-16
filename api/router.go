package api

import (
	"vidlearn-final-projcect/api/middleware"
	authV1 "vidlearn-final-projcect/api/v1/auth"
	categoryV1 "vidlearn-final-projcect/api/v1/category"
	playlistV1 "vidlearn-final-projcect/api/v1/playlist"
	playlistToolV1 "vidlearn-final-projcect/api/v1/playlist_tool"
	toolV1 "vidlearn-final-projcect/api/v1/tool"
	userV1 "vidlearn-final-projcect/api/v1/user"
	videoV1 "vidlearn-final-projcect/api/v1/video"
	"vidlearn-final-projcect/config"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	CategoryV1Controller     *categoryV1.Controller
	ToolV1Controller         *toolV1.Controller
	PlaylistV1Controller     *playlistV1.Controller
	VideoV1Controller        *videoV1.Controller
	PlaylistToolV1Controller *playlistToolV1.Controller
	UserV1Controller         *userV1.Controller
	AuthV1Controller         *authV1.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller, config *config.AppConfig) {
	categoryV1 := e.Group("/api/v1/category")
	categoryV1.GET("", controller.CategoryV1Controller.GetAll, middleware.JWTMiddleware(config))
	categoryV1.GET("/:id", controller.CategoryV1Controller.GetByID, middleware.JWTMiddleware(config))
	categoryV1.POST("", controller.CategoryV1Controller.Create, middleware.JWTMiddlewareAdmin(config))
	categoryV1.PUT("/:id", controller.CategoryV1Controller.Update, middleware.JWTMiddlewareAdmin(config))
	categoryV1.DELETE("/:id", controller.CategoryV1Controller.Delete, middleware.JWTMiddlewareAdmin(config))

	toolV1 := e.Group("/api/v1/tool")
	toolV1.GET("", controller.ToolV1Controller.GetAll, middleware.JWTMiddleware(config))
	toolV1.GET("/:id", controller.ToolV1Controller.GetByID, middleware.JWTMiddleware(config))
	toolV1.POST("", controller.ToolV1Controller.Create, middleware.JWTMiddlewareAdmin(config))
	toolV1.PUT("/:id", controller.ToolV1Controller.Update, middleware.JWTMiddlewareAdmin(config))
	toolV1.DELETE("/:id", controller.ToolV1Controller.Delete, middleware.JWTMiddlewareAdmin(config))

	playlistV1 := e.Group("/api/v1/playlist")
	playlistV1.GET("", controller.PlaylistV1Controller.GetAll, middleware.JWTMiddleware(config))
	playlistV1.GET("/:id", controller.PlaylistV1Controller.GetByID, middleware.JWTMiddleware(config))
	playlistV1.POST("", controller.PlaylistV1Controller.Create, middleware.JWTMiddlewareAdmin(config))
	playlistV1.PUT("/:id", controller.PlaylistV1Controller.Update, middleware.JWTMiddlewareAdmin(config))
	playlistV1.DELETE("/:id", controller.PlaylistV1Controller.Delete, middleware.JWTMiddlewareAdmin(config))

	videoV1 := e.Group("/api/v1/video")
	videoV1.GET("/:id/playlist", controller.VideoV1Controller.GetVideoByPlaylist, middleware.JWTMiddleware(config))
	videoV1.GET("/:id", controller.VideoV1Controller.GetByID, middleware.JWTMiddleware(config))
	videoV1.POST("", controller.VideoV1Controller.Create, middleware.JWTMiddlewareAdmin(config))
	videoV1.PUT("/:id", controller.VideoV1Controller.Update, middleware.JWTMiddlewareAdmin(config))
	videoV1.DELETE("/:id", controller.VideoV1Controller.Delete, middleware.JWTMiddlewareAdmin(config))

	playlistToolV1 := e.Group("/api/v1")
	playlistToolV1.GET("/playlist/:id/tool", controller.PlaylistToolV1Controller.GetAllToolsByPlaylist, middleware.JWTMiddleware(config))
	playlistToolV1.DELETE("/playlist/:id_playlist/tool/:id_tool", controller.PlaylistToolV1Controller.DeletePlaylistTool, middleware.JWTMiddlewareAdmin(config))
	playlistToolV1.POST("/playlist/:id/tool", controller.PlaylistToolV1Controller.CreatePlaylistTool, middleware.JWTMiddlewareAdmin(config))
	playlistToolV1.GET("/tool/:id/playlist", controller.PlaylistToolV1Controller.GetAllPlaylistsByTool, middleware.JWTMiddleware(config))

	userV1 := e.Group("/api/v1/user")
	userV1.GET("", controller.UserV1Controller.GetAll, middleware.JWTMiddlewareAdmin(config))
	userV1.GET("/:id", controller.UserV1Controller.GetByID, middleware.JWTMiddleware(config))
	userV1.PUT("/:id", controller.UserV1Controller.Update, middleware.JWTMiddlewareAdmin(config))
	userV1.DELETE("/:id", controller.UserV1Controller.Delete, middleware.JWTMiddlewareAdmin(config))

	authV1 := e.Group("/api/v1/auth")
	authV1.POST("/register", controller.AuthV1Controller.Register)
	authV1.POST("/login", controller.AuthV1Controller.Login)
	authV1.GET("/profile", controller.AuthV1Controller.GetProfile, middleware.JWTMiddleware(config))
	authV1.PUT("/profile", controller.AuthV1Controller.UpdateProfile, middleware.JWTMiddleware(config))
	authV1.PUT("/profile/change-password", controller.AuthV1Controller.ChangePassword, middleware.JWTMiddleware(config))

	verifyV1 := e.Group("/api/v1/verify")
	verifyV1.GET("/:token", controller.AuthV1Controller.Verify)

	forgotPasswordV1 := e.Group("/api/v1/forgot-password")
	forgotPasswordV1.POST("", controller.AuthV1Controller.ForgotPassword)

	resetPasswordV1 := e.Group("/api/v1/reset-password")
	resetPasswordV1.POST("/:token", controller.AuthV1Controller.ResetPassword)
}
