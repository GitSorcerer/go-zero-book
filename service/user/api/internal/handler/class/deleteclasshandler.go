package class

import (
	"net/http"

	"github.com/gonhon/go-zero-book/service/user/api/internal/logic/class"
	"github.com/gonhon/go-zero-book/service/user/api/internal/svc"
	"github.com/gonhon/go-zero-book/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteClassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClassPath
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := class.NewDeleteClassLogic(r.Context(), svcCtx)
		resp, err := l.DeleteClass(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
