package handler

import (
	"net/http"

	"github.com/fivezjd/go_zero_how/internal/logic"
	"github.com/fivezjd/go_zero_how/internal/svc"
	"github.com/fivezjd/go_zero_how/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Go_zero_howHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGo_zero_howLogic(r.Context(), svcCtx)
		resp, err := l.Go_zero_how(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
