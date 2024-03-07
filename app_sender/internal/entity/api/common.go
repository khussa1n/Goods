package api

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type RemoveGoods struct {
	Id      int64 `json:"id"`
	Removed bool  `json:"removed"`
}

type PayloadNewPriority struct {
	NewPriority int64 `json:"newPriority" binding:"required"`
}
