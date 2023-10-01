package main

import (
	"net/http"
	"task-5-pbi-btpns-febrian-syahroni/config"

	productController "task-5-pbi-btpns-febrian-syahroni/controllers/productController"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

var currentImage *imageupload.Image

func main() {
	r := gin.Default()
	config.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.GET("/image", func(c *gin.Context) {
		if currentImage == nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		currentImage.Write(c.Writer)
	})

	r.GET("/thumbnail", func(c *gin.Context) {
		if currentImage == nil {
			c.AbortWithStatus(http.StatusNotFound)
		}

		t, err := imageupload.ThumbnailJPEG(currentImage, 300, 300, 80)

		if err != nil {
			panic(err)
		}

		t.Write(c.Writer)
	})

	r.POST("/upload", func(c *gin.Context) {
		img, err := imageupload.Process(c.Request, "file")
		if err != nil {
			panic(err)
		}

		currentImage = img

		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.GET("/api/products", productController.Index)
	r.GET("/api/product/:id", productController.Show)
	r.POST("/api/product", productController.Create)
	r.PUT("/api/product/:id", productController.Update)
	r.DELETE("/api/product", productController.Delete)

	r.Run()
}