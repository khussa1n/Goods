package api

import (
	"github.com/khussa1n/Goods/app_sender/internal/entity"
)

type MetaGoods struct {
	Total   int64 `json:"total"`
	Removed int64 `json:"removed"`
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
}

type MetaProjects struct {
	Total  int64 `json:"total"`
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type GoodsList struct {
	Meta  MetaGoods      `json:"meta"`
	Goods []entity.Goods `json:"goods"`
}

type ProjectsList struct {
	Meta     MetaProjects      `json:"meta"`
	Projects []entity.Projects `json:"projects"`
}

type Priorities struct {
	Id       int64 `json:"id"`
	Priotiry int64 `json:"priotiry"`
}
