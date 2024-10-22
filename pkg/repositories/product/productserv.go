package product

import (
	"fmt"

	"github.com/sleepiinuts/simple-inventory-BE/pkg/models"
)

type ProductServ struct {
	repos ProductRepos
}

func (p *ProductServ) GetAll() ([]models.Product, error) {
	rows, err := p.repos.getAll()
	if err != nil {
		return nil, fmt.Errorf("product-serv-getAll [repos]: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product

		err = rows.StructScan(&product)
		if err != nil {
			return nil, fmt.Errorf("product-serv-getAll [scan]: %w", err)
		}

		// call RawToJson: Department
		err = product.RawToJson("Department")
		if err != nil {
			return nil, fmt.Errorf("product-serv-getAll [rawToJson]: %w", err)
		}

		products = append(products, product)
	}
	return products, nil
}

type Paginator struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
	Total     int `json:"total"`
}
type PagingProduct struct {
	Products []models.Product `json:"products"`
	Page     Paginator        `json:"page"`
}

func (p *ProductServ) GetAllWithPaging(pageIndex, pageSize int) (PagingProduct, error) {
	var pp PagingProduct

	// return paginator value
	pp.Page.PageIndex = pageIndex
	pp.Page.PageSize = pageSize

	row, err := p.repos.countAll()
	if err != nil {
		return pp, fmt.Errorf("product-serv-getAllWithPaging [repos]: %w", err)
	}

	if err = row.Scan(&pp.Page.Total); err != nil {
		return pp, fmt.Errorf("product-serv-getAllWithPaging [get total]: %w", err)
	}

	rows, err := p.repos.getAllWithPaging(pageIndex*pageSize+1, (pageIndex+1)*pageSize)
	if err != nil {
		return pp, fmt.Errorf("product-serv-getAllWithPaging [repos]: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product

		err = rows.StructScan(&product)
		if err != nil {
			return pp, fmt.Errorf("product-serv-getAllWithPaging [scan]: %w", err)
		}

		// call RawToJson: Department
		err = product.RawToJson("Department")
		if err != nil {
			return pp, fmt.Errorf("product-serv-getAll [rawToJson]: %w", err)
		}

		pp.Products = append(pp.Products, product)
	}

	return pp, nil
}

func NewProductServ(repos ProductRepos) *ProductServ {
	return &ProductServ{repos: repos}
}
