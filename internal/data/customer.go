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

	Insert(Customer) (Customer, error)
	Update(Customer) (Customer, error)
	Delete(id int64) error

	Page(pageParams pgdb.OffsetPageParams) CustomersQ

	FilterByID(ids ...int64) CustomersQ
	FilterByPersonID(ids ...int64) CustomersQ
	FilterByUserID(ids ...int64) CustomersQ
	FilterByDateBefore(time time.Time) CustomersQ
	FilterByDateAfter(time time.Time) CustomersQ

	JoinPerson() CustomersQ
}

type Customer struct {
	ID               int64      `db:"id" structs:"-"`
	PersonID         int64      `db:"person_id" structs:"person_id"`
	UserID           int64      `db:"user_id" structs:"user_id"`
	RegistrationDate *time.Time `db:"registration_date" structs:"registration_date"`
}
