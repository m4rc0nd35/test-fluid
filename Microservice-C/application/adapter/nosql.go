package adapter

type DatabaseNoSqlAdapter interface {
	Save(database, collect string, data any) (string, error)
}
