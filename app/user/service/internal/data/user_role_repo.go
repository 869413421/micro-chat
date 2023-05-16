package data

import (
	"fmt"
	"github.com/869413421/micro-chat/app/user/service/internal/biz"
	"github.com/casbin/casbin/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-kratos/kratos/v2/log"
)

// UserRoleRepo 数据库操作层
type userRoleRepo struct {
	data     *Data
	enforcer *casbin.Enforcer
	log      *log.Helper
}

// NewUserRoleRepo 新建DAO操作仓库
func NewUserRoleRepo(data *Data, logger log.Logger) biz.UserRoleRepo {
	return &userRoleRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user_role")),
	}
}

// Get 根据条件获取
func (r *userRoleRepo) Get(where map[string]interface{}) (*biz.UserRole, error) {
	var dbUserRole UserRole
	db := r.data.db
	for key, value := range where {
		db = db.Where(fmt.Sprintf("%s ?", key), value)
	}

	result := r.data.db.Where(where).First(&dbUserRole)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户角色不存在")
	}
	return &dbUserRole, nil
}
