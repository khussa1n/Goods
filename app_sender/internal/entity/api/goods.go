package api

type GoodsReq struct {
	Name        string `json:"name" db:"name" binding:"required"`
	Description string `json:"description" db:"description" binding:"required"`
}
