package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gitlab.com/zharzhanov/myguru/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.MaxMultipartMemory = 100 << 20

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		courses := api.Group("/courses")
		{
			courses.GET("", h.getCourses)
			courses.GET("/trending", h.getTrendingCourses)
			courses.GET("/soon", h.getComingSoonCourses)
			courses.GET("/:id", h.getCourse)
		}

		mentors := api.Group("/mentors")
		{
			mentors.GET("", h.getMentors)
			mentors.GET("/:id", h.getMentor)
		}
	}

	return router
}