package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/qustavo/dotsql"
)

func prepSqlLoader() {
	dots = make(map[string]*dotsql.DotSql)
	basePath, _ := os.Getwd()
	basePath = filepath.Dir(basePath) + "/pkg/repositories"

	if dot, err := dotsql.LoadFromFile(basePath + "/product/product.sql"); true {
		if err != nil {
			log.Fatal("Product sql loader error: ", err)
		}

		dots["product"] = dot
	}

}
