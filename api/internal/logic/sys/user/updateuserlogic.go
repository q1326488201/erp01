package user

import (
	"app/api/internal/svc"
	"app/api/internal/types"
	"app/common/errorf"
	"app/dao/models"
	"app/dao/query"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateAuthUserReq) (resp *types.ResultResponse, err error) {
	_, err = query.AuthUser.Where(query.AuthUser.ID.Eq(req.Id)).Updates(models.AuthUser{
		Username:  req.Username,
		Password:  req.Password,
		Nickname:  req.Nickname,
		AvatarURL: req.Avatar_url,
		Ftype:     req.Ftype,
		Fnote:     req.Fnote,
	})
	if err != nil {
		return nil, errorf.NewDefaultError(err.Error())
	}
	return types.NewDefaultResponse(), nil
}
