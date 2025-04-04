package controllers

import (
	"SparkGuardBackend/cmd/rest/controllers/adoptions"
	"SparkGuardBackend/cmd/rest/controllers/docs"
	"SparkGuardBackend/cmd/rest/controllers/events"
	"SparkGuardBackend/cmd/rest/controllers/groups"
	"SparkGuardBackend/cmd/rest/controllers/runner"
	"SparkGuardBackend/cmd/rest/controllers/students"
	"SparkGuardBackend/cmd/rest/controllers/tasks"
	"SparkGuardBackend/cmd/rest/controllers/users"
	"SparkGuardBackend/cmd/rest/controllers/work"
	"SparkGuardBackend/cmd/rest/middleware"
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

	r.Use(middleware.AuthMiddleware)
	students.SetupControllers(r)
	groups.SetupControllers(r)
	events.SetupControllers(r)
	work.SetupControllers(r)
	runner.SetupControllers(r)
	tasks.SetupControllers(r)
	adoptions.SetupControllers(r)

	return r
}
