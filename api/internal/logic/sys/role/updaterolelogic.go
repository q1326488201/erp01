package role

import (
	"app/common/errorf"
	"app/dao/models"
	"app/dao/query"
	"context"

	"app/api/internal/svc"
	"app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateAuthRoleReq) (resp *types.ResultResponse, err error) {
	_, err = query.AuthRole.Where(query.AuthRole.ID.Eq(req.Id)).Updates(models.AuthRole{
		Fname: req.Fname,
		Fnote: req.Fnote,
	})
	if err != nil {
		return nil, errorf.NewDefaultError(err.Error())
	}
	return types.NewDefaultResponse(), nil
}
