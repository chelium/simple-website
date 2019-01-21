package main

import (
	"path"
	"path/filepath"

	"github.com/chelium/simple-website/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./web/dist/web/index.html")
		} else {
			c.File("./web/dist/web/" + path.Join(dir, file))
		}
	})

	r.GET("/todo", handlers.GetTodoHandler)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
