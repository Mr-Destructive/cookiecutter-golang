package main

import (
  "github.com/gin-gonic/gin"
  "{{ cookiecutter.project_name }}/models"
)

func main() {
  r := gin.Default()
  models.ConnectDatabase()
  r.Run()
}
