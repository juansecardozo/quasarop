package infrastructures

import (
	"database/sql"

	"github.com/juansecardozo/quasarop/interfaces"
)

type PostgresHandler struct {
	Conn *sql.DB
}

func (handler *PostgresHandler) Execute(statement string, dest ...interface{}) error {
	_, err := handler.Conn.Exec(statement, dest...)

	if err != nil {
		return err
	}

	return nil
}

func (handler *PostgresHandler) Query(statement string, dest ...interface{}) (interfaces.IRow, error) {
	rows, err := handler.Conn.Query(statement, dest...)

	if err != nil {
		return new(PgRow), err
	}

	row := new(PgRow)
	row.Rows = rows

	return row, nil
}

type PgRow struct {
	Rows *sql.Rows
}

func (r PgRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)

	if err != nil {
		return err
	}

	return nil
}

func (r PgRow) Next() bool {
	return r.Rows.Next()
}
