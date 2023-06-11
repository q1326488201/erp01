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

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *AddUserLogic) AddUser(req *types.AddAuthUserReq) (resp *types.ResultResponse, err error) {
	count, err := query.AuthUser.Where(query.AuthUser.Username.Eq(req.Username)).Count()
	if err != nil {
		return nil, errorf.NewDefaultError(err.Error())
	}
	if count != 0 {
		return nil, errorf.NewDefaultError("账号已存在")
	}
	if err = query.AuthUser.Create(&models.AuthUser{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
		Fnote:    req.Fnote,
		Ftype:    int64(req.Ftype),
	}); err != nil {
		return nil, errorf.NewDefaultError(err.Error())
	}
	return types.NewDefaultResponse(), nil
}
