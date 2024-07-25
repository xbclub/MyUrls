package logic

import (
	"context"
	"github.com/xbclub/MyUrls/xerr"

	"github.com/xbclub/MyUrls/internal/svc"
	"github.com/xbclub/MyUrls/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortToLongLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortToLongLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortToLongLogic {
	return &ShortToLongLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortToLongLogic) ShortToLong(req *types.ShortToLongHandlerRequest) (resp *types.ShortToLongHandlerResponse, err error) {
	get, err := l.svcCtx.RedisC.Get(redisKeyPrefix + req.ShortKey)
	if err != nil || get == "" {
		l.Logger.Errorf("failed to get long URL")
		return nil, xerr.NewEnsumError(xerr.ResponseCodeServerError)
	}
	return &types.ShortToLongHandlerResponse{
		LongUrl: get,
	}, nil
}
