package sqlite3_sqlboiler

import (
	"github.com/hdget/common/types"
	"go.uber.org/fx"
)

const (
	providerName = "sqlite3-sqlboiler"
)

var Capability = types.Capability{
	Name:     providerName,
	Category: types.ProviderCategoryDb,
	Module: fx.Module(
		providerName,
		fx.Provide(New),
	),
}
