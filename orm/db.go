package orm

import (
	"fmt"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"strings"
)

// DBType database type
type DBType string

const (
	// dbMySQL Gorm Drivers mysql || postgres || sqlite || sqlserver
	dbMySQL      DBType = "mysql"
	dbPostgres   DBType = "postgres"
	dbSQLite     DBType = "sqlite"
	dbSQLServer  DBType = "sqlserver"
	dbClickHouse DBType = "clickhouse"
)

// ConnectDB choose db type for connection to database
func ConnectDB(t string, dsn string) (gorm.Dialector, error) {
	if dsn == "" {
		return nil, fmt.Errorf("dsn cannot be empty")
	}

	switch DBType(strings.ToLower(t)) {
	case dbMySQL:
		return mysql.Open(dsn), nil
	case dbPostgres:
		return postgres.Open(dsn), nil
	case dbSQLite:
		return sqlite.Open(dsn), nil
	case dbSQLServer:
		return sqlserver.Open(dsn), nil
	case dbClickHouse:
		return clickhouse.Open(dsn), nil
	default:
		return nil, fmt.Errorf("unknow db %q (support mysql || postgres || sqlite || sqlserver for now)", t)
	}
}
