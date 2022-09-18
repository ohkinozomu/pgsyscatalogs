package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/catalog-pg-depend.html
type PgDepend struct {
	Classid     pgtype.OID
	Objid       pgtype.OID
	Objsubid    pgtype.Int4
	Refclassid  pgtype.OID
	Refobjid    pgtype.OID
	Refobjsubid pgtype.Int4
	Deptype     pgtype.QChar
}
