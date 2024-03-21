package logic

import (
	"context"
	"github.com/alice52/awesome/web/go-zero/user/api/internal/svc"
	"github.com/alice52/awesome/web/go-zero/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.UserReply, err error) {
	// todo: add your logic here and delete this line

	l.Logger.Info("sasad")

	return
}
