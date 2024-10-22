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
	return o.queryHelper("GetAll", nil)
}

// getAllWithPaging implements ProductRepos.
func (o *OracleProductRepos) getAllWithPaging(pageIndex int, pageSize int) (*sqlx.Rows, error) {
	return o.queryHelper("GetAllWithPaging", []any{pageIndex, pageSize})
}

func (o *OracleProductRepos) queryHelper(stmtName string, args []any) (*sqlx.Rows, error) {
	stmt, err := o.dot.Raw(stmtName)
	if err != nil {
		return nil, fmt.Errorf("product-oracleRepos[%s]: %w", stmtName, err)
	}

	return o.db.Queryx(stmt, args...)
}

// countAll implements ProductRepos.
func (o *OracleProductRepos) countAll() (*sqlx.Row, error) {
	stmtName := "CountAll"
	stmt, err := o.dot.Raw(stmtName)
	if err != nil {
		return nil, fmt.Errorf("product-oracleRepos[%s]: %w", stmtName, err)
	}

	return o.db.QueryRowx(stmt), nil
}

func NewOracleProductRepos(db *sqlx.DB, dot *dotsql.DotSql) *OracleProductRepos {
	return &OracleProductRepos{
		db: db, dot: dot,
	}
}

var _ ProductRepos = &OracleProductRepos{}
