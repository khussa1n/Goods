package service

import (
	"context"
	"fmt"
	"github.com/khussa1n/Goods/app_sender/internal/entity"
	"github.com/khussa1n/Goods/app_sender/internal/entity/api"
	"github.com/khussa1n/Goods/app_sender/internal/natspkg"
	"log"
	"time"
)

func (m *Manager) CreateGood(ctx context.Context, g *entity.Goods) (*entity.Goods, error) {
	g.CreatedAt = time.Now()
	g.Removed = false
	good, err := m.PostgresRepository.CreateGood(ctx, g)
	if err != nil {
		return nil, fmt.Errorf("can not create Goods: %w", err)
	}

	err = natspkg.Publish(m.NatsConn, m.Config.Nats.Topic, good)
	if err != nil {
		return nil, err
	}

	err = m.RedisRepository.InvalidateGoodsCache(ctx)
	if err != nil {
		return nil, err
	}

	return good, nil
}

func (m *Manager) GetAllGoods(ctx context.Context, limit int64, offset int64) (*api.GoodsList, error) {
	// данные из Redis
	goodsList, err := m.RedisRepository.GetGoodsFromCache(ctx)
	if err == nil {
		return goodsList, nil
	}

	total, removed, goods, err := m.PostgresRepository.GetAllGoods(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("can not get all Goods: %w", err)
	}

	meta := api.MetaGoods{
		Total:   total,
		Removed: removed,
		Limit:   limit,
		Offset:  offset,
	}

	NewGoodsList := &api.GoodsList{Meta: meta, Goods: goods}

	err = m.RedisRepository.CacheGoods(ctx, NewGoodsList)
	if err != nil {
		log.Printf("error caching goods in Redis: %v", err)
	}

	return NewGoodsList, nil
}

func (m *Manager) DeleteGoodByID(ctx context.Context, id int64) error {
	good, err := m.PostgresRepository.DeleteGoodByID(ctx, id)
	if err != nil {
		return err
	}

	err = natspkg.Publish(m.NatsConn, m.Config.Nats.Topic, good)
	if err != nil {
		return err
	}

	err = m.RedisRepository.InvalidateGoodsCache(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) UpdateGoodByID(ctx context.Context, id int64, g *entity.Goods) (*entity.Goods, error) {
	good, err := m.PostgresRepository.UpdateGoodByID(ctx, id, g)
	if err != nil {
		return nil, err
	}

	err = natspkg.Publish(m.NatsConn, m.Config.Nats.Topic, good)
	if err != nil {
		return nil, err
	}

	err = m.RedisRepository.InvalidateGoodsCache(ctx)
	if err != nil {
		return nil, err
	}

	return good, nil
}

func (m *Manager) Reprioritiize(ctx context.Context, id int64, p int64) ([]api.Priorities, error) {
	priorities, err := m.PostgresRepository.Reprioritiize(ctx, id, p)
	if err != nil {
		return nil, err
	}

	err = m.RedisRepository.InvalidateGoodsCache(ctx)
	if err != nil {
		return nil, err
	}

	return priorities, nil
}
