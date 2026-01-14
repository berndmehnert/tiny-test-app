package main

import (
	"fmt"
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

	markdown := []byte(fmt.Sprintf(`In a certain shop, everything is sold with a **20%% discount**.  
However, because of a new tax law, a **20%% sales tax** is then added **after** the discount.

Many people mistakenly believe the customer ends up paying the original price.

If the original price of an item is 1000 Euros,  
how much does the customer actually pay in the end,  
and how many dollars cheaper or more expensive is the final price compared to the original 1000 Euros?`))
	htmlContent := mdToHTML(markdown)
	policy := bluemonday.UGCPolicy()
	clean := policy.Sanitize(string(htmlContent))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "page.tmpl", gin.H{
			"Title":        "Tiny Site",
			"ProblemTitle": "For O. M., a problem with a surprising result!",
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
