package logic

import (
	"app/api/internal/svc"
	"app/api/internal/types"
	"app/dao/models"
	"app/dao/query"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test() (resp *types.ResultResponse, err error) {
	err = query.AuthUser.Create(&models.AuthUser{
		ID:       0,
		Username: "123",
		Password: "123",
		Nickname: "123",
		Fnote:    "123",
	})
	if err != nil {
		return nil, err
	}
	return
}
