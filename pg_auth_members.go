package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/catalog-pg-auth-members.html
type PgAuthMembers struct {
	RoleID      pgtype.OID
	Member      pgtype.OID
	Grantor     pgtype.OID
	AdminOption pgtype.Bool
}
