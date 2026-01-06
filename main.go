package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	r.SetFuncMap(template.FuncMap{})
	r.LoadHTMLGlob("templates/*.tmpl")

	markdown := []byte("An astronaut on Mars harvests **100 kg** of Martian potatoes.  \nThese potatoes are composed of **99% water** and 1% solid potato matter.  \nThe astronaut leaves the potatoes outside in the dry Martian sun for a few hours. During this time, some of the water evaporates.  \n  \nWhen she returns, she measures the water content again and finds that the potatoes are now **98% water**.\n**What is the new total weight of the potatoes?**")
	htmlContent := mdToHTML(markdown)
	policy := bluemonday.UGCPolicy()
	clean := policy.Sanitize(string(htmlContent))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "page.tmpl", gin.H{
			"Title":        "Tiny Site",
			"ProblemTitle": "Another cool problem for the great O. M. from his dad concerning Martian potatoes!",
			"Message":      template.HTML(clean),
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_ = r.Run(":" + port)
}

func mdToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
