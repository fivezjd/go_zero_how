package logic

import (
	"context"

	"github.com/fivezjd/go_zero_how/internal/svc"
	"github.com/fivezjd/go_zero_how/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Go_zero_howLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGo_zero_howLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Go_zero_howLogic {
	return &Go_zero_howLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Go_zero_howLogic) Go_zero_how(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
