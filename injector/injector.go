package injector

import (
	"ahsmha/notes/infra"
)

func InjectDB() infra.SqlHandler {
	sqlhandler := infra.NewSqlHandler()
	return *sqlhandler
}
