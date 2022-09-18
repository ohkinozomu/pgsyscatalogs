package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/catalog-pg-cast.html
type PgCast struct {
	OID         pgtype.OID
	CastSource  pgtype.OID
	CastTarget  pgtype.OID
	CastFunc    pgtype.OID
	CastContext pgtype.QChar
	CastMethod  pgtype.QChar
}
