package dto

type Tree struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" form:"name" binding:"required"`
	ShortUrl    string `json:"short_url" form:"short_url" binding:"required"`
	OriginalUrl string `json:"original_url" form:"original_url" binding:"required"`
	CreatedAt   string `json:"created_at"`
}
