package product

import "github.com/jmoiron/sqlx"

type ProductRepos interface {
	getAll() (*sqlx.Rows, error)
}
