package menu

import (
	"app/dao/models"
	"app/dao/query"
	"context"

	"app/api/internal/svc"
	"app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMenuLogic) DeleteMenu(req *types.DeleteAuthMenuReq) (resp *types.ResultResponse, err error) {
	for _, id := range req.Id {
		query.AuthMenu.Unscoped().Delete(&models.AuthMenu{
			ID: id,
		})
	}
	return types.NewDefaultResponse(), nil
}
