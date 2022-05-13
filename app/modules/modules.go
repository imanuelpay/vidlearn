package modules

import (
	"vidlearn-final-projcect/api"
	categoryV1Controller "vidlearn-final-projcect/api/v1/category"
	categoryService "vidlearn-final-projcect/business/category"

	toolV1Controller "vidlearn-final-projcect/api/v1/tool"
	toolService "vidlearn-final-projcect/business/tool"

	playlistV1Controller "vidlearn-final-projcect/api/v1/playlist"
	playlistService "vidlearn-final-projcect/business/playlist"

	videoV1Controller "vidlearn-final-projcect/api/v1/video"
	videoService "vidlearn-final-projcect/business/video"

	playlistToolV1Controller "vidlearn-final-projcect/api/v1/playlist_tool"
	playlistToolService "vidlearn-final-projcect/business/playlist_tool"

	userV1Controller "vidlearn-final-projcect/api/v1/user"
	userService "vidlearn-final-projcect/business/user"

	authV1Controller "vidlearn-final-projcect/api/v1/auth"

	"vidlearn-final-projcect/config"
	categoryRepository "vidlearn-final-projcect/repository/category"
	playlistRepository "vidlearn-final-projcect/repository/playlist"
	playlistToolRepository "vidlearn-final-projcect/repository/playlist_tool"
	toolRepository "vidlearn-final-projcect/repository/tool"
	userRepository "vidlearn-final-projcect/repository/user"
	videoRepository "vidlearn-final-projcect/repository/video"
	"vidlearn-final-projcect/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	categoryPermitRepository := categoryRepository.CategoryRepository(dbCon)
	categoryPremitService := categoryService.CreateService(categoryPermitRepository)
	categoryV1PremitController := categoryV1Controller.CreateController(categoryPremitService)

	toolPermitRepository := toolRepository.ToolRepository(dbCon)
	toolPremitService := toolService.CreateService(toolPermitRepository)
	toolV1PremitController := toolV1Controller.CreateController(toolPremitService)

	playlistPermitRepository := playlistRepository.PlaylistRepository(dbCon)
	playlistPremitService := playlistService.CreateService(playlistPermitRepository)
	playlistV1PremitController := playlistV1Controller.CreateController(playlistPremitService)

	videoPermitRepository := videoRepository.VideoRepository(dbCon)
	videoPremitService := videoService.CreateService(videoPermitRepository)
	videoV1PremitController := videoV1Controller.CreateController(videoPremitService)

	playlistToolPermitRepository := playlistToolRepository.PlaylistToolRepository(dbCon)
	playlistToolPremitService := playlistToolService.CreateService(playlistToolPermitRepository)
	playlistToolV1PremitController := playlistToolV1Controller.CreateController(playlistToolPremitService)

	userPermitRepository := userRepository.UserRepository(dbCon)
	userPremitService := userService.CreateService(userPermitRepository, config)
	userV1PremitController := userV1Controller.CreateController(userPremitService)

	authV1PremitController := authV1Controller.CreateController(userPremitService, config)

	controllers := api.Controller{
		CategoryV1Controller:     categoryV1PremitController,
		ToolV1Controller:         toolV1PremitController,
		PlaylistV1Controller:     playlistV1PremitController,
		VideoV1Controller:        videoV1PremitController,
		PlaylistToolV1Controller: playlistToolV1PremitController,
		UserV1Controller:         userV1PremitController,
		AuthV1Controller:         authV1PremitController,
	}

	return controllers
}
