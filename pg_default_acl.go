package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/catalog-pg-default-acl.html
type PgDefaultACL struct {
	OID             pgtype.OID
	Defaclrole      pgtype.OID
	defaclnamespace pgtype.OID
	Defaclobjtype   pgtype.QChar
	defaclacl       []pgtype.ACLItem
}
