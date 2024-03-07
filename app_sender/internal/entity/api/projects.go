package api

type ProjectsReq struct {
	Name string `json:"name" db:"name" binding:"required"`
}
