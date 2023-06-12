package user

import (
	"app/api/internal/svc"
	"app/api/internal/types"
	"app/common/errorf"
	"app/dao/models"
	"app/dao/query"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type FindUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *FindUserLogic) FindUser(req *types.FindAuthUserReq) (resp *types.ResultResponse, err error) {
	lstUser, err := query.AuthUser.Select(query.AuthUser.ID, query.AuthUser.Username, query.AuthUser.CreatedAt,
		query.AuthUser.Nickname, query.AuthUser.Ftype, query.AuthUser.Fnote).
		Where(query.AuthUser.Username.Like("%"+req.Username+"%"),
			query.AuthUser.Ftype.Eq(req.Ftype),
			query.AuthUser.Nickname.Like("%"+req.Nickname+"%")).Find()
	if err != nil {
		return nil, errorf.NewDefaultError(err.Error())
	}
	lst := make([]types.FindAuthUserResp, 0)
	for _, user := range lstUser {
		Roles := make([]models.AuthRole, 0)
		query.AuthUserRole.Select(query.AuthRole.ID, query.AuthRole.Fname).Where(query.AuthUserRole.UserID.Eq(user.ID)).
			LeftJoin(query.AuthRole, query.AuthRole.ID.EqCol(query.AuthUserRole.RoleID)).Scan(&Roles)
		lst = append(lst, types.FindAuthUserResp{
			Id:             user.ID,
			Username:       user.Username,
			Nickname:       user.Nickname,
			Ftype:          user.Ftype,
			Avatar_url:     user.AvatarURL,
			Fnote:          user.Fnote,
			CreateDatetime: user.CreatedAt.Format(time.DateTime),
			Roles:          Roles,
		})
	}
	return types.NewDataResponse(lst), nil
}
