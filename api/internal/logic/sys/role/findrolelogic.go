package role

import (
	"app/common/errorf"
	"app/dao/models"
	"app/dao/query"
	"context"
	"time"

	"app/api/internal/svc"
	"app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleLogic {
	return &FindRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindRoleLogic) FindRole(req *types.FindAuthRoleReq) (resp *types.ResultResponse, err error) {
	lstRole, err := query.AuthRole.Debug().Select(query.AuthRole.ID, query.AuthRole.Fname, query.AuthRole.Fnote,
		query.AuthRole.CreatedAt).Where(query.AuthRole.Fname.Like("%" + req.Fname + "%")).Find()
	if err != nil {
		return nil, errorf.NewDefaultError(err.Error())
	}
	lst := make([]types.FindAuthRoleResp, 0)
	for _, role := range lstRole {
		Users := make([]models.AuthUser, 0)
		query.AuthUserRole.Select(query.AuthUser.ID, query.AuthUser.Username, query.AuthUser.Nickname).
			Where(query.AuthUserRole.RoleID.Eq(role.ID)).
			LeftJoin(query.AuthUser, query.AuthUser.ID.EqCol(query.AuthUserRole.UserID)).Scan(&Users)
		lst = append(lst, types.FindAuthRoleResp{
			Id:             role.ID,
			Fname:          role.Fname,
			Fnote:          role.Fnote,
			CreateDatetime: role.CreatedAt.Format(time.DateTime),
			Users:          Users,
		})
	}
	return types.NewDataResponse(lst), nil
}
