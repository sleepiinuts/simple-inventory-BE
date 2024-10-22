package product

import "github.com/jmoiron/sqlx"

type ProductRepos interface {
	getAll() (*sqlx.Rows, error)
	getAllWithPaging(pageIndex, pageSize int) (*sqlx.Rows, error)
	countAll() (*sqlx.Row, error)
}
