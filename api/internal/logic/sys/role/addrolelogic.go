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

type AddRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRoleLogic) AddRole(req *types.AddAuthRoleReq) (resp *types.ResultResponse, err error) {
	count, err := query.AuthRole.Where(query.AuthRole.Fname.Eq(req.Fname)).Count()
	if err != nil {
		return nil, errorf.NewDefaultError(err.Error())
	}
	if count != 0 {
		return nil, errorf.NewDefaultError("角色已存在")
	}
	if err = query.AuthRole.Create(&models.AuthRole{
		Fname: req.Fname,
		Fnote: req.Fnote,
	}); err != nil {
		return nil, errorf.NewDefaultError(err.Error())
	}
	return types.NewDefaultResponse(), nil
}
