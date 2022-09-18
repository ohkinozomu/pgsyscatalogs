package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/catalog-pg-description.html
type PgDescription struct {
	Objoid      pgtype.OID
	Classoid    pgtype.OID
	Objsubid    pgtype.Int4
	Description pgtype.Text
}
