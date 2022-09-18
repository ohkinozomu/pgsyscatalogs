package pgsyscatalogs

import (
	"github.com/jackc/pgtype"
)

// https://www.postgresql.org/docs/current/catalog-pg-event-trigger.html
type PgEventTrigger struct {
	OID        pgtype.OID
	Evtname    pgtype.Name
	Evtevent   pgtype.Name
	Evtowner   pgtype.OID
	Evtdfoid   pgtype.OID
	Evtenabled pgtype.QChar
	Evttags    []pgtype.Text
}
