package data

import (
	slog "log"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/869413421/micro-chat/app/user/service/internal/conf"
	schema2 "github.com/869413421/micro-chat/app/user/service/internal/data/orm/schema"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewUserRepo, NewRoleRepo, NewPermissionRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// migrate 模型迁移
func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&schema2.User{}, &schema2.UserRole{}, &schema2.Role{}, &schema2.RolePermission{}, &schema2.Permission{})
	if err != nil {
		return err
	}

	return nil
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	err := migrate(db)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

// NewDB .
func NewDB(c *conf.Data) (*gorm.DB, error) {
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢查询SQL阀值
			Colorful:      false,       // 禁止彩色打印
			LogLevel:      logger.Info, // 日等等级
		},
	)

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
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
