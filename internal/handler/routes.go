// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/fivezjd/go_zero_how/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: Go_zero_howHandler(serverCtx),
			},
		},
	)
}