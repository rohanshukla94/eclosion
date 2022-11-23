package eclosion

import "database/sql"

type appPaths struct {
	rootPath string
	dirNames []string
}
type databaseConfig struct {
	dsn      string
	database string
}

type Database struct {
	DataType string
	Pool     *sql.DB
}
