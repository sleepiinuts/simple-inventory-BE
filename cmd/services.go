package main

import "github.com/sleepiinuts/simple-inventory-BE/pkg/repositories/product"

func prepServices() {
	prodServ = product.NewProductServ(product.NewOracleProductRepos(db, dots["product"]))
}
