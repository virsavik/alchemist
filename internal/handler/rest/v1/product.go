package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"io.github.virsavik/alchemist/internal/model"
	"io.github.virsavik/alchemist/pkg/common/pagination"
)

func (h Handler) GetProducts(ctx *gin.Context) error {
	products, total, err := h.productCtrl.All(ctx, pagination.Input{
		Index:     1,
		Size:      2,
		WithTotal: true,
	})
	if err != nil {
		return err
	}

	page := pagination.Page[model.Product]{
		Items: products,
		Index: 1,
		Size:  2,
		Total: total,
	}

	ctx.JSON(http.StatusOK, page)

	return nil
}
