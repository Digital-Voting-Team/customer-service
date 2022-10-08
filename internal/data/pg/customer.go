package pg

import (
	"customer-service/internal/data"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"time"
)

const customersTableName = "public.customer"

func NewCustomersQ(db *pgdb.DB) data.CustomersQ {
	return &customersQ{
		db:        db.Clone(),
		sql:       sq.Select("customers.*").From(customersTableName),
		sqlUpdate: sq.Update(customersTableName).Suffix("returning *"),
	}
}

type customersQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlUpdate sq.UpdateBuilder
}

func (c *customersQ) New() data.CustomersQ {
	return NewCustomersQ(c.db)
}

func (c *customersQ) Get() (*data.Customer, error) {
	var result data.Customer
	err := c.db.Get(&result, c.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (c *customersQ) Select() ([]data.Customer, error) {
	var result []data.Customer
	err := c.db.Select(&result, c.sql)
	return result, err
}

func (c *customersQ) Transaction(fn func(q data.CustomersQ) error) error {
	return c.db.Transaction(func() error {
		return fn(c)
	})
}

func (c *customersQ) Insert(customer data.Customer) (data.Customer, error) {
	clauses := structs.Map(customer)
	clauses["person_id"] = customer.PersonID
	clauses["created_at"] = customer.CreatedAt

	var result data.Customer
	stmt := sq.Insert(customersTableName).SetMap(clauses).Suffix("returning *")
	err := c.db.Get(&result, stmt)

	return result, err
}

func (c *customersQ) Update(customer data.Customer) (data.Customer, error) {
	var result data.Customer
	clauses := structs.Map(customer)
	clauses["person_id"] = customer.PersonID
	clauses["created_at"] = customer.CreatedAt

	err := c.db.Get(&result, c.sqlUpdate.SetMap(clauses))
	return result, err
}

func (c *customersQ) Delete(id int64) error {
	stmt := sq.Delete(customersTableName).Where(sq.Eq{"id": id})
	err := c.db.Exec(stmt)
	return err
}

func (c *customersQ) Page(pageParams pgdb.OffsetPageParams) data.CustomersQ {
	c.sql = pageParams.ApplyTo(c.sql, "id")
	return c
}

func (c *customersQ) FilterByID(ids ...int64) data.CustomersQ {
	c.sql = c.sql.Where(sq.Eq{"id": ids})
	return c
}

func (c *customersQ) FilterByDateBefore(time time.Time) data.CustomersQ {
	stmt := sq.LtOrEq{"customer.created_at": time}
	c.sql = c.sql.Where(stmt)
	// Will not work for update
	// c.sqlUpdate = c.sqlUpdate.Where(stmt)
	return c
}

func (c *customersQ) FilterByDateAfter(time time.Time) data.CustomersQ {
	stmt := sq.GtOrEq{"customer.created_at": time}
	c.sql = c.sql.Where(stmt)
	// Will not work for update
	// c.sqlUpdate = c.sqlUpdate.Where(stmt)
	return c
}

func (c *customersQ) JoinPerson() data.CustomersQ {
	stmt := fmt.Sprintf("%s as customer on public.person.id = customer.person_id",
		customersTableName)
	c.sql = c.sql.Join(stmt)
	return c
}
