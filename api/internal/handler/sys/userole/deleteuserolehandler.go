package userole

import (
	"net/http"

	"app/api/internal/logic/sys/userole"
	"app/api/internal/svc"
	"app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteUseroleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userole.NewDeleteUseroleLogic(r.Context(), svcCtx)
		resp, err := l.DeleteUserole(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
