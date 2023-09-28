package userV2

import (
	"context"

	"github.com/fivezjd/go_zero_how/internal/svc"
	"github.com/fivezjd/go_zero_how/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterV2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterV2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterV2Logic {
	return &UserRegisterV2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterV2Logic) UserRegisterV2(req *types.RequestOne) (resp *types.ResponseOne, err error) {
	// todo: add your logic here and delete this line

	return
}
