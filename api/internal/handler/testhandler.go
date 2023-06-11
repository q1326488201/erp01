package handler

import (
	"net/http"

	"app/api/internal/logic"
	"app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func testHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewTestLogic(r.Context(), svcCtx)
		resp, err := l.Test()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
