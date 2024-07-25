package handler

import (
	"net/http"

	"github.com/xbclub/MyUrls/internal/logic"
	"github.com/xbclub/MyUrls/internal/svc"
	"github.com/xbclub/MyUrls/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LongToShortHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LongToShortHandlerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLongToShortLogic(r.Context(), svcCtx)
		resp, err := l.LongToShort(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
