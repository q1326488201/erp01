package user

import (
	"app/api/internal/svc"
	"app/api/internal/types"
	"app/dao/models"
	"app/dao/query"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteAuthUserReq) (resp *types.ResultResponse, err error) {
	for _, id := range req.Id {
		fmt.Println(id)
		query.AuthUser.Unscoped().Delete(&models.AuthUser{
			ID: id,
		})
	}
	return types.NewDefaultResponse(), nil
}
