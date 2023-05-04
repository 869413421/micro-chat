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
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewUserRepo, NewRoleRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// migrate 模型迁移
func migrate(db *gorm.DB) error {
	//var err error
	//// 手动创建唯一索引，因为gorm会重复创建唯一索引
	//if !db.Migrator().HasIndex(&User{}, "email") {
	//	err = db.Migrator().CreateIndex(&User{}, "email,unique")
	//	if err != nil {
	//		fmt.Printf("create user email index error: %v", err)
	//		os.Exit(1)
	//	}
	//}
	//
	//if !db.Migrator().HasIndex(&Role{}, "name") {
	//	err = db.Migrator().CreateIndex(&Role{}, "name,unique")
	//	if err != nil {
	//		fmt.Printf("create user name index error: %v", err)
	//	}
	//}
	err := db.AutoMigrate(&User{}, &UserRole{}, &Role{}, &RolePermission{}, &Permission{})
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

func NewDB(c *conf.Data) *gorm.DB {
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
		log.Errorf("connection to database error: %v", err)
	}

	return db
}
