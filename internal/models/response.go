package models

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//// BeforeCreate hook to generate a UUID before saving a new record
//func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
//	m.ID = uuid.New()
//	return nil
//}

type APIResponse struct {
	Success    bool         `json:"success"`
	Data       interface{}  `json:"data"`
	Error      *ErrorDetail `json:"error,omitempty"`
	Pagination *Pagination  `json:"pagination,omitempty"`
	Meta       ResponseMeta `json:"meta"`
}

type ErrorDetail struct {
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Details    interface{} `json:"details,omitempty"`
	HTTPStatus int         `json:"http_status"`
}

type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
	HasNext    bool  `json:"has_next"`
	HasPrev    bool  `json:"has_prev"`
}

type ResponseMeta struct {
	Timestamp string `json:"timestamp"`
	RequestID string `json:"request_id"`
	Version   string `json:"version"`
}
