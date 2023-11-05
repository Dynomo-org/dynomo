package handler

import "time"

type AppContent struct {
	ID           string    `json:"id"`
	AppID        string    `json:"app_id"`
	Title        string    `json:"title"`
	CategoryID   string    `json:"category_id"`
	Description  string    `json:"description"`
	Content      string    `json:"content"`
	ThumbnailURL string    `json:"thumbnail_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
