package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/catalog-pg-extension.html
type PgExtension struct {
	OID            pgtype.OID
	ExtName        pgtype.Name
	ExtOwner       pgtype.OID
	ExtNamespace   pgtype.OID
	ExtRelocatable pgtype.Bool
	ExtVersion     pgtype.Text
	ExtConfig      []pgtype.OID
	ExtCondition   []pgtype.Text
}
