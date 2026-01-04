package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	r.SetFuncMap(template.FuncMap{})
	r.LoadHTMLGlob("templates/*.tmpl")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "page.tmpl", gin.H{
			"Title":        "Tiny Site",
			"ProblemTitle": "A cool problem for the big O. M. from his dad concerning muffins!",
			"Message":      "A school bake sale sells muffins. At the start there are 120 muffins. During the morning they sell 40% of the muffins. After a break they sell another 25% of the muffins that were left after the morning. How many muffins remain at the end of the day?",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_ = r.Run(":" + port)
}
