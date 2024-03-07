package service

import (
	"context"
	"fmt"
	"github.com/khussa1n/Goods/app_sender/internal/entity"
	"github.com/khussa1n/Goods/app_sender/internal/entity/api"
	"time"
)

func (m *Manager) CreateProject(ctx context.Context, p *entity.Projects) (*entity.Projects, error) {
	p.CreatedAt = time.Now()
	project, err := m.PostgresRepository.CreateProject(ctx, p)
	if err != nil {
		return nil, fmt.Errorf("can not create Project: %w", err)
	}

	return project, nil
}

func (m *Manager) GetAllProjects(ctx context.Context, limit int64, offset int64) (*api.ProjectsList, error) {
	total, projects, err := m.PostgresRepository.GetAllProjects(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("can not get all Projects: %w", err)
	}

	meta := api.MetaProjects{
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}

	NewProjectsList := &api.ProjectsList{Meta: meta, Projects: projects}

	return NewProjectsList, nil
}

func (m *Manager) DeleteProjectByID(ctx context.Context, id int64) error {
	err := m.PostgresRepository.DeleteProjectByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) UpdateProjectByID(ctx context.Context, id int64, name string) (*entity.Projects, error) {
	project, err := m.PostgresRepository.UpdateProjectByID(ctx, id, name)
	if err != nil {
		return nil, err
	}

	return project, nil
}
