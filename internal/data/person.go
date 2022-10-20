package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"time"
)

type PersonsQ interface {
	New() PersonsQ

	Get() (*Person, error)
	Select() ([]Person, error)

	Transaction(fn func(q PersonsQ) error) error

	Insert(Person) (Person, error)
	Update(Person) (Person, error)
	Delete(id int64) error

	Page(pageParams pgdb.OffsetPageParams) PersonsQ

	FilterByID(ids ...int64) PersonsQ
	FilterByNames(names ...string) PersonsQ
	FilterByPhones(phones ...string) PersonsQ
	FilterByEmails(emails ...string) PersonsQ
	FilterByBirthday(date time.Time) PersonsQ

	JoinAddress() PersonsQ
}

type Person struct {
	ID        int64      `db:"id" structs:"-"`
	Name      string     `db:"name" structs:"name"`
	Phone     string     `db:"phone" structs:"phone"`
	Email     string     `db:"email" structs:"email"`
	Birthday  *time.Time `db:"birthday" structs:"birthday"`
	AddressID int64      `db:"address_id" structs:"address_id"`
}
