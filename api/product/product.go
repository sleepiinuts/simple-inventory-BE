package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sleepiinuts/simple-inventory-BE/pkg/repositories/product"
)

type ProductApi struct {
	prodServ *product.ProductServ
}

func (p *ProductApi) GetAll(c *gin.Context) {
	prods, err := p.prodServ.GetAll()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, prods)
}

func NewProductApi(prodServ *product.ProductServ) *ProductApi {
	return &ProductApi{prodServ: prodServ}
}
