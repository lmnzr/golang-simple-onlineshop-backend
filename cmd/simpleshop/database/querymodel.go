package database

import (
	"database/sql"

	"github.com/lmnzr/simpleshop/cmd/simpleshop/database/filter"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/database/group"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/database/order"
)

//QueryModel : Database Access Object
type QueryModel struct {
	DBConn        *sql.DB
	Model         interface{}
	Table         string
	Limit         int
	Offset        int
	Filters       []filter.Filter
	Orders        []order.Order
	Groups        []group.Group
	Error         []string
	DBConnTx      *sql.Tx
	Transaction   bool
}

//NewQuery : Create Database Access Object With Generated Query
func NewQuery(dbcon *sql.DB, tablename string, model interface{}) QueryModel {
	return QueryModel{
		DBConn: dbcon,
		Table:  tablename,
		Model:  model,
	}
}

//NewQueryCustom : Create Database Access Object With Custom Query
func NewQueryCustom(dbcon *sql.DB) QueryModel {
	return QueryModel{
		DBConn: dbcon,
	}
}	

//NewTransaction : Create Database Access Object With Generated Query (Transaction)
func NewTransaction(tx *sql.Tx, tablename string, model interface{}) QueryModel {
	return  QueryModel{
		Table:  tablename,
		Model:  model,
		Transaction: true,
		DBConnTx: tx,
	}
}

//NewTransactionCustom : Create Database Access Object With Custom Query (Transaction)
func NewTransactionCustom(tx *sql.Tx) QueryModel {
	return QueryModel{
		Transaction: true,
		DBConnTx: tx,
	}
}

//SetLimit : Set Query LIMIT only for SELECT
func (qm *QueryModel) SetLimit(limit int) *QueryModel {
	qm.Limit = limit
	return qm
}

//SetOffset : Set Query OFSSET if LIMIT defined only for SELECT
func (qm *QueryModel) SetOffset(offset int) *QueryModel {
	qm.Offset = offset
	return qm
}

//SetFilters : Set WHERE filter in Query
func (qm *QueryModel) SetFilters(filters []filter.Filter) *QueryModel {
	qm.Filters = filters
	return qm
}

//SetGroups : Set GROUP BY only for SELECT
func (qm *QueryModel) SetGroups(groups []group.Group) *QueryModel {
	qm.Groups = groups
	return qm
}

//SetOrders : Set ORDER BY only for SELECT
func (qm *QueryModel) SetOrders(orders []order.Order) *QueryModel {
	qm.Orders = orders
	return qm
}
