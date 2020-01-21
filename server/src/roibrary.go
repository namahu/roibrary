package main

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.HTMLRender = loadTemplates("./templates")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Html5 Template Engine",
		})
	})

	router.Run()

}

func loadTemplates(templatsDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatsDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatsDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}

	return r

}
