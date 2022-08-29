package api

import (
	//"{{ cookiecutter.project_name }}/models"
  "net/http"
	"github.com/gin-gonic/gin"
)
func Hello_World(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"data":"hello world"})
}
