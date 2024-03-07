package redisrepo

import (
	"context"
	"github.com/khussa1n/Goods/app_sender/internal/entity/api"
)

type Good interface {
	GetGoodsFromCache(ctx context.Context) (*api.GoodsList, error)
	CacheGoods(ctx context.Context, p *api.GoodsList) error
	InvalidateGoodsCache(ctx context.Context) error
}

type RedisRepository interface {
	Good
}
