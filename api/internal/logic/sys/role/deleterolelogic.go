package role

import (
	"app/dao/models"
	"app/dao/query"
	"context"
	"fmt"

	"app/api/internal/svc"
	"app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRoleLogic) DeleteRole(req *types.DeleteAuthRoleReq) (resp *types.ResultResponse, err error) {
	for _, id := range req.Id {
		fmt.Println(id)
		query.AuthRole.Unscoped().Delete(&models.AuthRole{
			ID: id,
		})
	}
	return types.NewDefaultResponse(), nil

}
