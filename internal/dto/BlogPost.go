package dto

import "time"

type BlogpostDTO struct {
	BlogpostID string     `json:"blogpostId"`
	Title      string     `json:"title,omitempty"`
	Markdown   string     `json:"markdown,omitempty"`
	Category   string     `json:"category,omitempty"`
	Image      string     `json:"image,omitempty"`
	Video      string     `json:"video,omitempty"`
	Date       *time.Time `json:"date,omitempty"`
	Published  bool       `json:"published"`
}
