package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/catalog-pg-foreign-server.html
type PgForeignServer struct {
	OID        pgtype.OID
	Srvname    pgtype.Name
	Srvowner   pgtype.OID
	Srvfdw     pgtype.OID
	Srvtype    pgtype.Text
	Srvversion pgtype.Text
	Srvacl     []pgtype.ACLItem
	Srvoptions []pgtype.Text
}
