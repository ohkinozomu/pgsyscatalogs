package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/catalog-pg-database.html
type PgDatabase struct {
	OID           pgtype.OID
	Datname       pgtype.Name
	Datdba        pgtype.OID
	Encoding      pgtype.Int4
	Datcollate    pgtype.Name
	Datctype      pgtype.Name
	Datistemplate pgtype.Bool
	Datallowconn  pgtype.Bool
	Datconnlimit  pgtype.Int4
	Datlastsysoid pgtype.OID
	Datfrozenxid  pgtype.XID
	Datminmxid    pgtype.XID
	Dattablespace pgtype.OID
	Datacl        []pgtype.ACLItem
}
