package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/shiki-tak/shiki-web/server/interfaces/database"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	DB *sql.DB
	Tx *sql.Tx
}

func NewSqlHandler() *SqlHandler {
	// "user:password@tcp(localhost:3306)/database"
	dns := fmt.Sprintf("%s:%s@%s(%s)/%s", "user", "password", "tcp", "localhost:3306", "database")
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatalf("Could not open db: %v", err)
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.DB = db
	return sqlHandler
}

func (handler *SqlHandler) TxBegin() error {
	tx, err := handler.DB.Begin()
	if err != nil {
		return err
	}

	handler.Tx = tx
	return nil
}

func (handler *SqlHandler) TxCommit() error {
	if err := handler.Tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (handler *SqlHandler) TxRollback() error {
	if err := handler.Tx.Rollback(); err != nil {
		return err
	}
	return nil
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Tx.Exec(statement, args...)
	if err != nil {
		return res, err
	}

	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	row := new(SqlRow)
	rows, err := handler.DB.Query(statement, args...)
	if err != nil {
		return row, err
	}
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r *SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r *SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r *SqlRow) Close() error {
	return r.Rows.Close()
}
