package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"io.github.virsavik/alchemist/internal/controller/product"
)

type Handler struct {
	productCtrl product.Controller
}

func New(productCtrl product.Controller) Handler {
	return Handler{
		productCtrl: productCtrl,
	}
}

func HandleEndpoint(f func(ctx *gin.Context) error) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if err := f(ctx); err != nil {
			log.Printf("request error: %s\n", err)

			if herr, ok := err.(WebErr); ok {
				ctx.JSON(herr.Status, gin.H{
					"message": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "something went wrong",
			})
		}
	}
}
