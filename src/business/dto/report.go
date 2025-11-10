package dto

import (
	"mime/multipart"
	"time"

	"github.com/nupalHariz/LaporIn/src/business/entity"
)

type InputReport struct {
	Title       string               `form:"title" binding:"required"`
	Description string               `form:"description" binding:"required"`
	Category    string               `form:"category" binding:"required"`
	Location    string               `form:"location" binding:"required"`
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
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

type ReportParam struct {
	Id int `uri:"id"`
	entity.PaginationParam
}

type GetReport struct {
	Id             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Category       string    `json:"category"`
	Location       string    `json:"location"`
	PhotoUrl       string    `json:"photo_url"`
	TicketCode     string    `json:"ticket_code"`
	Status         string    `json:"status"`
	StatusDesc     string    `json:"status_desc"`
	StatusProofUrl string    `json:"status_proof_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UpdateParam struct {
	Id              int                  `uri:"id"`
	Status          string               `form:"status" binding:"required"`
	StatusDesc      string               `form:"status_desc"`
	StatusProofFile multipart.FileHeader `form:"status_proof_file"`
}
