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

func NewProductServ(repos ProductRepos) *ProductServ {
	return &ProductServ{repos: repos}
}
