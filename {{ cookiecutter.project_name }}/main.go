package main

import (
  "github.com/gin-gonic/gin"
  "{{ cookiecutter.project_name }}/models"
  "{{ cookiecutter.project_name }}/views"
)

func main() {
  r := gin.Default()
  r.LoadHTMLGlob("templates/**")
  models.ConnectDatabase()
  r.GET("/", api.Hello_World)
  r.GET("/hello", api.Render_Hello)
  r.Run()
}
