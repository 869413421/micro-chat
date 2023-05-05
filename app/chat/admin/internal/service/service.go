package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	v1 "github.com/869413421/micro-chat/api/chat/admin/v1"
	"github.com/869413421/micro-chat/app/chat/admin/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewChatAdmin)

var _ v1.ChatAdminServer = (*ChatAdmin)(nil)

type ChatAdmin struct {
	v1.UnimplementedChatAdminServer
	log *log.Helper
	uc  biz.UserUsecase
	rc  biz.RoleUsecase
}

// NewChatAdmin 创建用户服务。
func NewChatAdmin(uc biz.UserUsecase, rc biz.RoleUsecase, logger log.Logger) *ChatAdmin {
	return &ChatAdmin{
		uc:  uc,
		rc:  rc,
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
	}
}
