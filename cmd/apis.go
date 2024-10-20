package main

import productApi "github.com/sleepiinuts/simple-inventory-BE/api/product"

func prepApis() {
	prodApi = productApi.NewProductApi(prodServ)
}
