package userole

import (
	"app/api/internal/svc"
	"app/api/internal/types"
	"app/common/errorf"
	"app/dao/models"
	"app/dao/query"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUseroleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUseroleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUseroleLogic {
	return &AddUseroleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUseroleLogic) AddUserole(req *types.UserRoleReq) (resp *types.ResultResponse, err error) {
	if len(req.UserRoleLst) == 0 {
		return nil, errorf.NewDefaultError("请选择角色或用户")
	}
	for _, userRole := range req.UserRoleLst {
		if userRole.UserId > 0 && userRole.RoleId > 0 {
			if c, err := query.AuthUserRole.Where(query.AuthUserRole.UserID.Eq(userRole.UserId),
				query.AuthUserRole.RoleID.Eq(userRole.RoleId)).Count(); err != nil {
				return nil, errorf.NewDefaultError(err.Error())
			} else {
				if c > 0 {
					continue
				}
			}
			query.AuthUserRole.Create(&models.AuthUserRole{
				UserID: userRole.UserId,
				RoleID: userRole.RoleId,
			})
		}
	}
	return types.NewDefaultResponse(), nil
}
