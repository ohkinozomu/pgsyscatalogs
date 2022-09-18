package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/catalog-pg-collation.html
type PgCollation struct {
	OID                 pgtype.OID
	CollName            pgtype.Name
	CollNamespace       pgtype.OID
	CollOwner           pgtype.OID
	CollProvider        pgtype.QChar
	CollIsdeterministic pgtype.Bool
	CollEncoding        pgtype.Int4
	CollCollate         pgtype.Name
	CollCtype           pgtype.Name
	CollVersion         pgtype.Text
}
