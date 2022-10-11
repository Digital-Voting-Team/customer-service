package pg

import (
	"customer-service/internal/data"
	"database/sql"
	"gitlab.com/distributed_lab/kit/pgdb"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
)

const addressesTableName = "public.address"

func NewAddressesQ(db *pgdb.DB) data.AddressesQ {
	return &addressesQ{
		db:        db.Clone(),
		sql:       sq.Select("address.*").From(addressesTableName),
		sqlUpdate: sq.Update(addressesTableName).Suffix("returning *"),
	}
}

type addressesQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlUpdate sq.UpdateBuilder
}

func (q *addressesQ) New() data.AddressesQ {
	return NewAddressesQ(q.db)
}

func (q *addressesQ) Get() (*data.Address, error) {
	var result data.Address
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *addressesQ) Select() ([]data.Address, error) {
	var result []data.Address
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *addressesQ) Update(address data.Address) (data.Address, error) {
	var result data.Address
	clauses := structs.Map(address)
	clauses["building_number"] = address.BuildingNumber
	clauses["street"] = address.Street
	clauses["city"] = address.City
	clauses["district"] = address.District
	clauses["region"] = address.Region
	clauses["postal_code"] = address.PostalCode

	err := q.db.Get(&result, q.sqlUpdate.SetMap(clauses))

	return result, err
}

func (q *addressesQ) Transaction(fn func(q data.AddressesQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *addressesQ) Insert(address data.Address) (data.Address, error) {
	clauses := structs.Map(address)
	clauses["building_number"] = address.BuildingNumber
	clauses["street"] = address.Street
	clauses["city"] = address.City
	clauses["district"] = address.District
	clauses["region"] = address.Region
	clauses["postal_code"] = address.PostalCode

	var result data.Address
	stmt := sq.Insert(addressesTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return result, err
}

func (q *addressesQ) Delete(id int64) error {
	stmt := sq.Delete(addressesTableName).Where(sq.Eq{"id": id})
	err := q.db.Exec(stmt)
	return err
}

func (q *addressesQ) Page(pageParams pgdb.OffsetPageParams) data.AddressesQ {
	q.sql = pageParams.ApplyTo(q.sql, "id")
	return q
}

func (q *addressesQ) FilterByID(ids ...int64) data.AddressesQ {
	q.sql = q.sql.Where(sq.Eq{"id": ids})
	q.sqlUpdate = q.sqlUpdate.Where(sq.Eq{"id": ids})
	return q
}

func (q *addressesQ) FilterByBuildingNumber(numbers ...int64) data.AddressesQ {
	q.sql = q.sql.Where(sq.Eq{"building_number": numbers})
	return q
}

func (q *addressesQ) FilterByStreet(streets ...string) data.AddressesQ {
	q.sql = q.sql.Where(sq.Eq{"street": streets})
	return q
}

func (q *addressesQ) FilterByCities(cities ...string) data.AddressesQ {
	q.sql = q.sql.Where(sq.Eq{"city": cities})
	return q
}

func (q *addressesQ) FilterByDistricts(districts ...string) data.AddressesQ {
	q.sql = q.sql.Where(sq.Eq{"district": districts})
	return q
}

func (q *addressesQ) FilterByRegion(regions ...string) data.AddressesQ {
	q.sql = q.sql.Where(sq.Eq{"region": regions})
	return q
}

func (q *addressesQ) FilterByPostalCodes(codes ...string) data.AddressesQ {
	q.sql = q.sql.Where(sq.Eq{"postal_code": codes})
	return q
}
