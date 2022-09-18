package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/catalog-pg-authid.html
type PgAuthid struct {
	OID            pgtype.OID
	RolName        pgtype.Bool
	RolSuper       pgtype.Bool
	RolInherit     pgtype.Bool
	RolCreaterole  pgtype.Bool
	RolCreatedb    pgtype.Bool
	RolCanlogin    pgtype.Bool
	RolReplication pgtype.Bool
	RolBypassrls   pgtype.Bool
	RolConnlimit   pgtype.Int4
	RolPassword    pgtype.Text
	RolValiduntil  pgtype.Timestamptz
}
