package menu

import (
	"app/api/internal/svc"
	"app/api/internal/types"
	"app/common/errorf"
	"app/dao/query"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAuthMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindAuthMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAuthMenuLogic {
	return &FindAuthMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *FindAuthMenuLogic) FindAuthMenu(req *types.FindAuthMenuReq) (resp *types.ResultResponse, err error) {
	lstMenu, err := query.AuthMenu.Where(query.AuthMenu.Name.Eq(req.FName)).Find()
	if err != nil {
		return nil, errorf.NewDefaultError(err.Error())
	}
	lst := make([]types.FindAuthMenuResp, 0)
	for _, menu := range lstMenu {
		lst = append(lst, types.FindAuthMenuResp{
			FID:           menu.ID,
			FName:         menu.Name,
			FParentId:     menu.ParentID,
			FUrl:          menu.URL,
			FPerms:        menu.Perms,
			FType:         menu.Ftype,
			FIcon:         menu.Icon,
			FOrderNum:     menu.OrderNum,
			FVuePath:      menu.VuePath,
			FVueComponent: menu.VueComponent,
			FVueRedirect:  menu.VueRedirect,
		})
	}
	return types.NewDataResponse(lst), nil
}
