package main

import (
	"go-todo/domain"
	"go-todo/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*.html")

	db := infrastructure.DbInit()
	defer db.Close()
	
	router.GET("/", func(c *gin.Context) {
		todos := infrastructure.DbRead()
		c.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})
	
	router.POST("/new", func(c *gin.Context) {
		text := c.PostForm("text")   // フォームの値を取得できる。取得できる値の型はstring
		rawStatus := c.PostForm("status")
		id, err := strconv.Atoi(rawStatus)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		status := domain.Status(id)
		deadline := c.PostForm("deadline")
		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		todo := domain.Todo{Text: text, Status: status, Deadline: deadline}
	
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		infrastructure.DbCreate(todo)
		c.Redirect(302, "/")
	})

  router.Run()
}
