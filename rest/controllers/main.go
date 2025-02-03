package controllers

import (
	"SparkGuardBackend/controllers/docs"
	"SparkGuardBackend/controllers/events"
	"SparkGuardBackend/controllers/groups"
	"SparkGuardBackend/controllers/students"
	"SparkGuardBackend/controllers/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupCORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}

func setupSwagger(r *gin.Engine) {
	docs.SwaggerInfo.Title = "SPARK GUARD API"
	docs.SwaggerInfo.Description = "This is a REST API for Spark Guard"
	docs.SwaggerInfo.BasePath = "/"

	swaggerUrl := ginSwagger.URL("./doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))
}

func SetupRouter() *gin.Engine {
	r := gin.New()

	setupCORS(r)
	setupSwagger(r)

	users.SetupControllers(r)
	students.SetupControllers(r)
	groups.SetupControllers(r)
	events.SetupControllers(r)

	return r
}
