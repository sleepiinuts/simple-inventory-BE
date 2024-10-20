package main

import (
	"github.com/jmoiron/sqlx"
	go_ora "github.com/sijms/go-ora/v2"
)

func conn() *sqlx.DB {
	username := "SYS"
	password := "myP@ssw0rd"
	host := "localhost"
	port := 1521
	serviceName := "LPMDB"

	connStr := go_ora.BuildUrl(host, port, serviceName, username, password, nil)
	db, err := sqlx.Open("oracle", connStr)

	if err != nil {
		logger.Error("Connection Error: ", "error", err)
	} else {
		logger.Info("Database connected!")
	}

	return db
}
