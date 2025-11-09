package dto

import (
	"mime/multipart"
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
