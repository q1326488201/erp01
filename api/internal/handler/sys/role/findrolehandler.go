package role

import (
	"net/http"

	"app/api/internal/logic/sys/role"
	"app/api/internal/svc"
	"app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FindRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindAuthRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewFindRoleLogic(r.Context(), svcCtx)
		resp, err := l.FindRole(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
