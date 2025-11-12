package entity

import (
	"time"

	"github.com/reyhanmichiels/go-pkg/v2/null"
	"github.com/reyhanmichiels/go-pkg/v2/query"
)

type Status string
type Category string

const (
	INFRASTRUCTURE Category = "Infrastuktur"
	SERVICE        Category = "Pelayanan"
	SECURITY       Category = "Keamanan"

	NEW      Status = "New"
	INREVIEW Status = "In Review"
	REJECTED Status = "Rejected"
	RESOLVED Status = "Resolved"
)

type Report struct {
	Id             int         `db:"id"`
	Title          string      `db:"title"`
	Description    string      `db:"description"`
	Category       Category    `db:"category"`
	Location       string      `db:"location"`
	PhotoUrl       null.String `db:"photo_url"`
	TicketCode     string      `db:"ticket_code"`
	Status         Status      `db:"status"`
	StatusDesc     null.String `db:"status_desc"`
	StatusProofUrl null.String `db:"status_proof_url"`
	CreatedAt      time.Time   `db:"created_at"`
	UpdatedAt      time.Time   `db:"updated_at"`
}

type ReportInputParam struct {
	Title       string      `db:"title"`
	Description string      `db:"description"`
	Category    Category    `db:"category"`
	Location    string      `db:"location"`
	PhotoUrl    null.String `db:"photo_url"`
	TicketCode  string      `db:"ticket_code"`
}

type ReportParam struct {
	Id              int       `db:"id" param:"id"`
	Title           string    `db:"title" param:"LIKE"`
	Category        Category  `db:"category" param:"category"`
	Status          Status    `db:"status" param:"status"`
	CreatedAtHigher time.Time `db:"created_at" param:"created_at__gte"`
	CreatedAtLower  time.Time `db:"created_at" param:"created_at__lte"`
	Option          query.Option
	PaginationParam
}

type UpdateReportParam struct {
	Status         Status      `db:"status"`
	StatusDesc     null.String `db:"status_desc"`
	StatusProofUrl null.String `db:"status_proof_url"`
}
