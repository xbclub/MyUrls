package logic

import (
	"context"
	"encoding/base64"
	"github.com/xbclub/MyUrls/xerr"
	"math/rand"
	"time"

	"github.com/xbclub/MyUrls/internal/svc"
	"github.com/xbclub/MyUrls/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const redisKeyPrefix = "myurls:"

type LongToShortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLongToShortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LongToShortLogic {
	return &LongToShortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LongToShortLogic) LongToShort(req *types.LongToShortHandlerRequest) (resp *types.LongToShortHandlerResponse, err error) {
	longUrl := req.LongUrl
	shortKey := req.ShortKey
	// 兼容以前的实现，这里如果是 base64 编码的字符串，进行解码
	_longUrl, err := base64.StdEncoding.DecodeString(longUrl)
	if err == nil {
		longUrl = string(_longUrl)
	}
	// generate short key
	if shortKey == "" {
		shortKey = generateRandomString(l.svcCtx.Config.ShortKeyLength)
	}
	exists, err := l.svcCtx.RedisC.ExistsCtx(l.ctx, redisKeyPrefix+shortKey)
	if err != nil {
		l.Logger.Errorf("failed to check short key redis error: %v", err)
		return nil, xerr.NewEnsumError(xerr.ResponseCodeServerError)
	}
	if exists {
		return nil, xerr.NewEnsumError(xerr.ResponseCodeParamsCheckError)
	}
	err = l.svcCtx.RedisC.SetexCtx(l.ctx, redisKeyPrefix+shortKey, longUrl, l.svcCtx.Config.ShortKeyTTL)
	if err != nil {
		l.Logger.Errorf("failed to create short URL redis error: %v", err)
		return nil, xerr.NewEnsumError(xerr.ResponseCodeServerError)
	}
	resp = &types.LongToShortHandlerResponse{
		Code:     xerr.ResponseCodeSuccessLegacy,
		ShortUrl: l.svcCtx.Config.WebSiteURL + "/" + shortKey,
	}
	return
}
func generateRandomString(bits int) string {
	// Create a byte slice b of length bits.
	b := make([]byte, bits)

	// Create a new random number generator with the current time as the seed.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random byte for each element in the byte slice b using the letterBytes slice.
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}

	// Convert the byte slice to a string and return it.
	return string(b)
}
