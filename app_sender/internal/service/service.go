package service

import (
	"context"
	"github.com/khussa1n/Goods/app_sender/internal/entity"
	"github.com/khussa1n/Goods/app_sender/internal/entity/api"
)

type Project interface {
	CreateProject(ctx context.Context, p *entity.Projects) (*entity.Projects, error)
	GetAllProjects(ctx context.Context, limit int64, offset int64) (*api.ProjectsList, error)
	DeleteProjectByID(ctx context.Context, id int64) error
	UpdateProjectByID(ctx context.Context, id int64, name string) (*entity.Projects, error)
}

type Good interface {
	CreateGood(ctx context.Context, g *entity.Goods) (*entity.Goods, error)
	GetAllGoods(ctx context.Context, limit int64, offset int64) (*api.GoodsList, error)
	DeleteGoodByID(ctx context.Context, id int64) error
	UpdateGoodByID(ctx context.Context, id int64, g *entity.Goods) (*entity.Goods, error)
	Reprioritiize(ctx context.Context, id int64, p int64) ([]api.Priorities, error)
}

type Service interface {
	Project
	Good
}
