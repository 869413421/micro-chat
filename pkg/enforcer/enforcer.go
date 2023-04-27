package enforcer

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewEnforcer)

func NewEnforcer(db *gorm.DB) (*casbin.Enforcer, error) {
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

	m, err := model.NewModelFromString(`
        [request_definition]
        r = obj, act

        [policy_definition]
        p = obj, act, eft

        [policy_effect]
        e = some(where (p.eft == allow))

        [matchers]
        m = r.obj == p.obj && r.act == p.act
    `)
	if err != nil {
		return nil, err
	}

	return casbin.NewEnforcer(m, a)
}
