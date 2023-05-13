package base

import "github.com/volatiletech/sqlboiler/v4/queries/qm"

type QueryInput interface {
	ToQueryMods() []qm.QueryMod
}
