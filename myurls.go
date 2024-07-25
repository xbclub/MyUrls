package main

import (
	"flag"
	"fmt"
	"github.com/xbclub/MyUrls/internal/config"
	"github.com/xbclub/MyUrls/internal/handler"
	"github.com/xbclub/MyUrls/internal/svc"
	"github.com/xbclub/MyUrls/xerr"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	"go/types"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	xhttp "github.com/zeromicro/x/http"
)

var configFile = flag.String("f", "etc/myurls.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	defer server.Stop()
	httpx.SetErrorHandler(func(err error) (int, any) {
		switch e := err.(type) {
		case *errors.CodeMsg:
			return http.StatusOK, xhttp.BaseResponse[types.Nil]{
				Code: e.Code,
				Msg:  e.Msg,
			}
		default:
			return http.StatusOK, xhttp.BaseResponse[types.Nil]{
				Code: xerr.ResponseCodeServerError,
				Msg:  e.Error(),
			}
		}
	})
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
