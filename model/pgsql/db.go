package pgsql

import (
	"banwire/dash/dashboard_front/db"
	"banwire/dash/dashboard_front/model"

	pg "github.com/jackc/pgx"
)

// Db stores and manages the database connection
type Db struct {
	*pg.Conn
}

// GetConnection() returns the connection (*pg.Conn)
func (d *Db) GetConnection() *pg.Conn {
	return d.Conn
}

// ViewRequest handels the parameters of the ViewRequest object for postgres
type ViewRequest struct {
	*model.ViewRequest
}

// WhereJson returns the json string of 'condition' parameter
func (v *ViewRequest) WhereJson() interface{} {
	return db.ToPgJson(v.Condition)
}

// OrderJson returns the json string of 'order' parameter
func (v *ViewRequest) OrderJson() interface{} {
	return db.ToPgJson(v.Order)
}

// Error is a set of parameters that represents an error object
type Error struct {
	Type    string
	Message string
	Source  string
}

// PgdbObj returns the parameters of error object in a slice interface for the scan
func (o *Error) PgdbObj() []interface{} {
	return []interface{}{&o.Type, &o.Message, &o.Source}
}

// Error returns the string message of the error object
func (o *Error) Error() string {
	return o.Message
}
