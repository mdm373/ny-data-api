package db

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Masterminds/structable"

	_ "github.com/lib/pq" //postgres driver needed for sql connection
	"github.com/pkg/errors"
)

//Connection stateful connection to our db
type Connection struct {
	db *sql.DB
}

func (c *Connection) RunQuery(builder squirrel.SelectBuilder) (*sql.Rows, error) {
	return builder.PlaceholderFormat(squirrel.Dollar).RunWith(c.db).Query()
}

func (c *Connection) Bind(tableName string, schema structable.Record) structable.Recorder {
	return structable.New(squirrel.NewStmtCacheProxy(c.db), "postgres").Bind(tableName, schema)
}

var EmptyConnection = Connection{}

//Get snags us a connection to our wonderful backend, or an error if something goes terribly wrong
func Get(user string, password string, host string) (Connection, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		host,
		5432,
		user,
		password,
		"postgres",
	))
	if err != nil {
		return EmptyConnection, errors.Wrap(err, "cannot connect to backend")
	}
	err = db.Ping()
	if err != nil {
		return EmptyConnection, errors.Wrap(err, "cannot ping backend")
	}
	return Connection{db: db}, nil
}
