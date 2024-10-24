package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sleepiinuts/simple-inventory-BE/pkg/repositories/product"
)

type ProductApi struct {
	prodServ *product.ProductServ
}

func (p *ProductApi) GetAll(c *gin.Context) {
	pageIndex, err1 := strconv.Atoi(c.Query("pageIndex"))
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
	if err1 == nil && err2 == nil {
		pp, err := p.prodServ.GetAllWithPaging(pageIndex, pageSize)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, pp)
		return
	}

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
