package injector

import (
	"ahsmha/Tasks/infra"
)

func InjectDB() infra.SqlHandler {
	sqlhandler := infra.NewSqlHandler()
	return *sqlhandler
}
