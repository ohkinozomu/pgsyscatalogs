package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/catalog-pg-db-role-setting.html
type PgDBRoleSetting struct {
	Setdatabase pgtype.OID
	Setrole     pgtype.OID
	Setconfig   []pgtype.Text
}
