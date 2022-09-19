package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/catalog-pg-foreign-data-wrapper.html
type PgForeignDataWrapper struct {
	OID          pgtype.OID
	FDWName      pgtype.Name
	FDWOwner     pgtype.OID
	FDWHandler   pgtype.OID
	FDWValidator pgtype.OID
	FDWACL       pgtype.OID
	FDWOptions   []pgtype.Text
}
