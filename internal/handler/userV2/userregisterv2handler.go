package userV2

import (
	"net/http"

	"github.com/fivezjd/go_zero_how/internal/logic/userV2"
	"github.com/fivezjd/go_zero_how/internal/svc"
	"github.com/fivezjd/go_zero_how/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserRegisterV2Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestOne
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userV2.NewUserRegisterV2Logic(r.Context(), svcCtx)
		resp, err := l.UserRegisterV2(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
