package redisrepo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/khussa1n/Goods/app_sender/internal/entity/api"
	"log"
	"time"
)

func (r *Redis) GetGoodsFromCache(ctx context.Context) (*api.GoodsList, error) {
	val, err := r.RedisClient.Get(ctx, GoodsList).Result()
	if err == redis.Nil {
		log.Printf("Goods list not found in cache\n")
		return nil, fmt.Errorf("goods list not found in cache")
	} else if err != nil {
		log.Printf("Error retrieving goods list from cache: %s\n", err)
		return nil, err
	}

	var goodsList api.GoodsList
	err = json.Unmarshal([]byte(val), &goodsList)
	if err != nil {
		log.Printf("Error unmarshalling goods list from cache: %v\n", err)
		return nil, fmt.Errorf("error unmarshalling goods list from cache: %v", err)
	}

	log.Printf("Retrieved goods list from cache successfully\n")
	return &goodsList, nil
}

func (r *Redis) CacheGoods(ctx context.Context, goodsList *api.GoodsList) error {
	val, err := json.Marshal(goodsList)
	if err != nil {
		log.Printf("Error marshalling goods list for cache: %v\n", err)
		return fmt.Errorf("error marshalling goods list for cache: %v", err)
	}

	expiration := 1 * time.Minute // кеширование на минуту
	err = r.RedisClient.SetEX(ctx, GoodsList, val, expiration).Err()
	if err != nil {
		log.Printf("Error caching goods list in Redis: %v\n", err)
		return fmt.Errorf("error caching goods list in Redis: %v", err)
	}

	log.Printf("Cached goods list in Redis successfully\n")
	return nil
}

func (r *Redis) InvalidateGoodsCache(ctx context.Context) error {
	err := r.RedisClient.Del(ctx, GoodsList).Err()
	if err != nil {
		log.Printf("Error invalidating goods cache in Redis: %v\n", err)
		return fmt.Errorf("error invalidating goods cache in Redis: %v", err)
	}

	log.Printf("Invalidated goods cache in Redis successfully\n")
	return nil
}
