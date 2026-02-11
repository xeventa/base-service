package public

import "database/sql"

// DBPinger abstracts database connectivity checks.
type DBPinger interface {
	Ping() error
}

// SQLDBPinger implements DBPinger using a standard *sql.DB.
type SQLDBPinger struct {
	DB *sql.DB
}

func (p SQLDBPinger) Ping() error {
	if p.DB == nil {
		return nil
	}
	return p.DB.Ping()
}
