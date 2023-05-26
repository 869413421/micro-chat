package enforcer

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	slog "log"
	"os"
	"time"
)

var enforce *casbin.SyncedEnforcer

// GetSyncedEnforcer 获取casbin的enforcer
func GetSyncedEnforcer() (*casbin.SyncedEnforcer, error) {
	if enforce == nil {
		db, err := GetAdapterDb()
		if err != nil {
			return nil, err
		}

		enforce, err = NewSyncedEnforcer(db)
		if err != nil {
			return nil, err
		}
	}
	return enforce, nil
}

// NewSyncedEnforcer 创建casbin的enforcer
func NewSyncedEnforcer(db *gorm.DB) (*casbin.SyncedEnforcer, error) {
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

	m, err := model.NewModelFromString(`
       	[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act || r.sub == "admin"
    `)
	if err != nil {
		return nil, err
	}

	return casbin.NewSyncedEnforcer(m, a)
}

// GetAdapterDb 获取gorm的db
func GetAdapterDb() (*gorm.DB, error) {
	source := "root:root@tcp(user-db:33061)/user?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢查询SQL阀值
			Colorful:      false,       // 禁止彩色打印
			LogLevel:      logger.Info, // 日等等级
		},
	)

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
