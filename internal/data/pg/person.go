package pg

import (
	"customer-service/internal/data"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const personsTableName = "public.person"

func NewPersonsQ(db *pgdb.DB) data.PersonsQ {
	return &personsQ{
		db:        db.Clone(),
		sql:       sq.Select("person.*").From(personsTableName),
		sqlUpdate: sq.Update(personsTableName).Suffix("returning *"),
	}
}

type personsQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlUpdate sq.UpdateBuilder
}

func (p *personsQ) New() data.PersonsQ {
	return NewPersonsQ(p.db)
}

func (p *personsQ) Get() (*data.Person, error) {
	var result data.Person
	err := p.db.Get(&result, p.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (p *personsQ) Select() ([]data.Person, error) {
	var result []data.Person
	err := p.db.Select(&result, p.sql)
	return result, err
}

func (p *personsQ) Update(person data.Person) (data.Person, error) {
	var result data.Person
	clauses := structs.Map(person)
	clauses["name"] = person.Name
	clauses["phone"] = person.Phone
	clauses["email"] = person.Email
	clauses["address_id"] = person.AddressID

	err := p.db.Get(&result, p.sqlUpdate.SetMap(clauses))
	return result, err
}

func (p *personsQ) Transaction(fn func(q data.PersonsQ) error) error {
	return p.db.Transaction(func() error {
		return fn(p)
	})
}

func (p *personsQ) Insert(person data.Person) (data.Person, error) {
	clauses := structs.Map(person)
	clauses["name"] = person.Name
	clauses["phone"] = person.Phone
	clauses["email"] = person.Email
	clauses["address_id"] = person.AddressID

	var result data.Person
	stmt := sq.Insert(personsTableName).SetMap(clauses).Suffix("returning *")
	err := p.db.Get(&result, stmt)

	return result, err
}

func (p *personsQ) Delete(id int64) error {
	stmt := sq.Delete(personsTableName).Where(sq.Eq{"id": id})
	err := p.db.Exec(stmt)
	return err
}

func (p *personsQ) Page(pageParams pgdb.OffsetPageParams) data.PersonsQ {
	p.sql = pageParams.ApplyTo(p.sql, "id")
	return p
}

func (p *personsQ) FilterByID(ids ...int64) data.PersonsQ {
	p.sql = p.sql.Where(sq.Eq{"id": ids})
	return p
}

func (p *personsQ) FilterByNames(names ...string) data.PersonsQ {
	p.sql = p.sql.Where(sq.Eq{"name": names})
	return p
}

func (p *personsQ) FilterByPhones(phones ...string) data.PersonsQ {
	p.sql = p.sql.Where(sq.Eq{"phone": phones})
	return p
}

func (p *personsQ) FilterByEmails(emails ...string) data.PersonsQ {
	p.sql = p.sql.Where(sq.Eq{"email": emails})
	return p
}

func (p *personsQ) JoinAddress() data.PersonsQ {
	stmt := fmt.Sprintf("%s as person on public.address.id = person.address_id",
		personsTableName)
	p.sql = p.sql.Join(stmt)
	return p
}
