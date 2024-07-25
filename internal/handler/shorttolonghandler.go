package handler

import (
	"net/http"

	"github.com/xbclub/MyUrls/internal/logic"
	"github.com/xbclub/MyUrls/internal/svc"
	"github.com/xbclub/MyUrls/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShortToLongHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShortToLongHandlerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShortToLongLogic(r.Context(), svcCtx)
		resp, err := l.ShortToLong(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			//httpx.OkJsonCtx(r.Context(), w, resp)
			http.Redirect(w, r, resp.LongUrl, http.StatusMovedPermanently)
		}
	}
}
