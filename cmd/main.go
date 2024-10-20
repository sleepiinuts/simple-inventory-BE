package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/qustavo/dotsql"
	productApi "github.com/sleepiinuts/simple-inventory-BE/api/product"
	"github.com/sleepiinuts/simple-inventory-BE/pkg/repositories/product"
)

var (
	logger   *slog.Logger
	db       *sqlx.DB
	dots     map[string]*dotsql.DotSql
	prodServ *product.ProductServ
	prodApi  *productApi.ProductApi
	router   *gin.Engine
)

func main() {
	// create logger
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Starting Simple-Inventory BE")
	// create db
	db = conn()
	defer db.Close()

	if err := db.Ping(); err != nil {
		logger.Error("PING error: ", "error", err)
	}

	// prep sql loader
	prepSqlLoader()

	// prep services
	prepServices()

	// prep apis
	prepApis()

	// prep router
	getRouter(logger)
}
