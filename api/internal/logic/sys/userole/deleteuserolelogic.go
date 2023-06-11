package userole

import (
	"app/common/errorf"
	"app/dao/query"
	"context"

	"app/api/internal/svc"
	"app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUseroleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUseroleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUseroleLogic {
	return &DeleteUseroleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUseroleLogic) DeleteUserole(req *types.UserRoleReq) (resp *types.ResultResponse, err error) {
	if len(req.UserRoleLst) == 0 {
		return nil, errorf.NewDefaultError("请选择角色或用户")
	}
	for _, userRole := range req.UserRoleLst {
		if userRole.UserId > 0 && userRole.RoleId > 0 {
			query.AuthUserRole.Unscoped().Where(query.AuthUserRole.UserID.Eq(userRole.UserId),
				query.AuthUserRole.RoleID.Eq(userRole.RoleId)).Delete()
		}
	}
	return types.NewDefaultResponse(), nil
}
