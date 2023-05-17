// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/869413421/micro-chat/app/user/service/internal/biz"
	"github.com/869413421/micro-chat/app/user/service/internal/conf"
	"github.com/869413421/micro-chat/app/user/service/internal/data"
	"github.com/869413421/micro-chat/app/user/service/internal/server"
	"github.com/869413421/micro-chat/app/user/service/internal/service"
	"github.com/869413421/micro-chat/pkg/enforcer"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	db, err := data.NewDB(confData)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup, err := data.NewData(confData, logger, db)
	if err != nil {
		return nil, nil, err
	}
	casbinEnforcer, err := enforcer.NewEnforcer(db)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, casbinEnforcer, logger)
	roleRepo := data.NewRoleRepo(dataData, casbinEnforcer, logger)
	userUsecase := biz.NewUserUsecase(userRepo, roleRepo, auth, logger)
	roleUsecase := biz.NewRoleUsecase(roleRepo, logger)
	permissionRepo := data.NewPermissionRepo(dataData, casbinEnforcer, logger)
	permissionUsecase := biz.NewPermissionUsecase(permissionRepo, logger)
	userService := service.NewUserService(userUsecase, roleUsecase, permissionUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
