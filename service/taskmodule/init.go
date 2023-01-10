package taskmodule

import (
	"database/sql"
)

type Module struct {
	Storage *storage
}

func NewTaskModule(db *sql.DB) *Module {
	return &Module{
		Storage: newStorage(db),
	}
}
