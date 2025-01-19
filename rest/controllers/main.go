package controllers

import (
	"SparkGuardBackend/controllers/users"
	"github.com/gin-gonic/gin"
)

func SetupControllers(r *gin.Engine) {
	users.SetupControllers(r)
}
