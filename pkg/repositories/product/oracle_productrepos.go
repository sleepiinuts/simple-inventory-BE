package product

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/qustavo/dotsql"
)

type OracleProductRepos struct {
	db  *sqlx.DB
	dot *dotsql.DotSql
}

// getAll implements ProductRepos.
func (o *OracleProductRepos) getAll() (*sqlx.Rows, error) {
	stmt, err := o.dot.Raw("GetAll")
	if err != nil {
		return nil, fmt.Errorf("product-oracleRepos: %w", err)
	}

	return o.db.Queryx(stmt)
}

func NewOracleProductRepos(db *sqlx.DB, dot *dotsql.DotSql) *OracleProductRepos {
	return &OracleProductRepos{
		db: db, dot: dot,
	}
}

var _ ProductRepos = &OracleProductRepos{}
