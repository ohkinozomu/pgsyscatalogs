package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/view-pg-tables.html
type PgTables struct {
	SchemaName  pgtype.Name
	TableName   pgtype.Name
	TableOwner  pgtype.Name
	TableSpace  pgtype.Name
	HasIndexes  pgtype.Bool
	HasRules    pgtype.Bool
	HasTriggers pgtype.Bool
	RoqSecurity pgtype.Bool
}
