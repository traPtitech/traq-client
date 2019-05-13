package traq

import "time"

type ErrorResponse struct {
	Code    int
	Message string `json:"message"`
}

func (er *ErrorResponse) Error() string {
	return er.Message
}

type MessageResponse struct {
	MessageID       string                 `json:"messageId"`
	UserID          string                 `json:"userId"`
	ParentChannelID string                 `json:"parentChannelId"`
	Content         string                 `json:"content"`
	CreatedAt       time.Time              `json:"createdAt"`
	UpdatedAt       time.Time              `json:"updatedAt"`
	Pin             bool                   `json:"pin"`
	Reported        bool                   `json:"reported"`
	StampList       []MessageStampResponse `json:"stampList"`
}

type MessageStampResponse struct {
	StampID   string    `json:"stampId"`
	UserID    string    `json:"userId"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
