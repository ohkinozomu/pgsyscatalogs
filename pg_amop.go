package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/catalog-pg-amop.html
type PgAmop struct {
	OID            pgtype.OID
	Amopfamily     pgtype.OID
	Amoplefttype   pgtype.OID
	Amoprighttype  pgtype.OID
	Amopstrategy   pgtype.Int2
	Amoppurpose    pgtype.QChar
	Amopopr        pgtype.OID
	Amopmethod     pgtype.OID
	Amopsortfamily pgtype.OID
}
