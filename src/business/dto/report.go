package dto

import (
	"mime/multipart"
	"time"
)

type InputReport struct {
	Title       string               `form:"title"`
	Description string               `form:"description"`
	Category    string               `form:"category"`
	Location    string               `form:"location"`
	PhotoFile   multipart.FileHeader `form:"photo_file"`
}

type InputReportResponse struct {
	TicketCode string `json:"ticket_code"`
}

type AllReports struct {
	Id         int       `json:"id"`
	TicketCode string    `json:"ticket_code"`
	Title      string    `json:"title"`
	Category   string    `json:"category"`
	Location   string    `json:"location"`
	CreatedAt  time.Time `json:"created_at"`
}
