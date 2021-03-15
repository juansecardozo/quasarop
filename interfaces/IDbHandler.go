package interfaces

type IDbHandler interface {
	Execute(statement string, dest ...interface{}) error
	Query(statement string, dest ...interface{}) (IRow, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}
