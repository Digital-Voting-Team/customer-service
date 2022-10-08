package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"time"
)

type CustomersQ interface {
	New() CustomersQ

	Get() (*Customer, error)
	Select() ([]Customer, error)

	Transaction(fn func(q CustomersQ) error) error

	Insert(customer Customer) (Customer, error)
	Update(customer Customer) ([]Customer, error)
	Delete(id int64) error

	Page(pageParams pgdb.OffsetPageParams) CustomersQ

	FilterByID(ids ...int64) CustomersQ
	FilterByDateBefore(time time.Time) CustomersQ
	FilterByDateAfter(time time.Time) CustomersQ

	JoinPerson() CustomersQ
}

type Customer struct {
	ID        int64      `db:"id" structs:"-"`
	PersonID  string     `db:"person_id" structs:"person_id"`
	CreatedAt *time.Time `db:"created_at" structs:"created_at"`
}
