package types

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

//DBContext :
type DBContext struct {
	echo.Context
	*sql.DB
}
