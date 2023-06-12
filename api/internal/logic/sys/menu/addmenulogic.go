package menu

import (
	"app/api/internal/svc"
	"app/api/internal/types"
	"app/common/errorf"
	"app/dao/models"
	"app/dao/query"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMenuLogic) AddMenu(req *types.AddAuthMenuReq) (resp *types.ResultResponse, err error) {
	if count, err := query.AuthMenu.Where(query.AuthMenu.Name.Eq(req.FName)).Count(); err != nil {
		return nil, errorf.NewDefaultError(err.Error())
	} else {
		if count > 0 {
			return nil, errorf.NewDefaultError("菜单已存在")
		}
	}
	if err := query.AuthMenu.Create(&models.AuthMenu{
		Name:         req.FName,
		ParentID:     req.FParentId,
		URL:          req.FUrl,
		Icon:         req.FIcon,
		OrderNum:     req.FOrderNum,
		VuePath:      req.FVuePath,
		VueComponent: req.FVueComponent,
		VueRedirect:  req.FVueRedirect,
		Perms:        req.FPerms,
	}); err != nil {
		return nil, errorf.NewDefaultError(err.Error())
	}
	return types.NewDefaultResponse(), nil
}
