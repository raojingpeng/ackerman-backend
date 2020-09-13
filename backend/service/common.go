package service

type PageInfo struct {
	Page     int `form:"page" binding:"required,gte=1"`
	PageSize int `form:"page_size" binding:"required,gte=5,lte=100"`
}
