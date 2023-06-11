package api

import (
	"app/api/internal/config"
	"app/api/internal/handler"
	"app/api/internal/svc"
	"app/common/errorf"
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

var configFile = flag.String("f", "./api/etc/api-api.yaml", "the config file")

func Start() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, any) {
		if err != nil {
			return http.StatusOK, errorf.NewCodeErrorResponse(err.Error())
		}
		return http.StatusOK, nil
	})
	httpx.SetErrorHandler(func(err error) (int, any) {
		switch err.(type) {
		case *errorf.CodeError:
			return http.StatusOK, err.(*errorf.CodeError).Data()
		default:
			return http.StatusOK, errorf.NewCodeErrorResponse(err.Error())
		}
	})
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
