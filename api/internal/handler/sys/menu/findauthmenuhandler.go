package menu

import (
	"net/http"

	"app/api/internal/logic/sys/menu"
	"app/api/internal/svc"
	"app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FindAuthMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindAuthMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewFindAuthMenuLogic(r.Context(), svcCtx)
		resp, err := l.FindAuthMenu(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
