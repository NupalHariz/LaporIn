package report

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"

	reportDomain "github.com/nupalHariz/LaporIn/src/business/domain/report"
	"github.com/nupalHariz/LaporIn/src/business/dto"
	"github.com/nupalHariz/LaporIn/src/business/entity"
	"github.com/nupalHariz/LaporIn/src/business/service/supabase"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
	"github.com/reyhanmichiels/go-pkg/v2/null"
)

type Interface interface {
	InputReport(ctx context.Context, inputParam dto.InputReport) (dto.InputReportResponse, error)
	GetAllReports(ctx context.Context, param dto.ReportParam) ([]dto.AllReports, error)
	GetReport(ctx context.Context, param dto.ReportParam) (dto.GetReport, error)
	UpdateReport(ctx context.Context, param dto.UpdateParam) error
}

type report struct {
	report   reportDomain.Interface
	supabase supabase.Interface
}

type InitParam struct {
	Report   reportDomain.Interface
	Supabase supabase.Interface
}

func Init(param InitParam) Interface {
	return &report{
		report:   param.Report,
		supabase: param.Supabase,
	}
}

var wib = time.FixedZone("WIB", 7*3600)

func (r *report) InputReport(ctx context.Context, param dto.InputReport) (dto.InputReportResponse, error) {
	res := dto.InputReportResponse{}

	ticketCode, err := r.GenerateTicketCode(time.Now(), wib)
	if err != nil {
		return res, err
	}

	var photoUrl string

	if param.PhotoFile.Size != 0 {
		now := time.Now().In(wib)

		key := fmt.Sprintf("%s/%d-%s", now.Format("02-01-2006"), now.UnixNano(), param.PhotoFile.Filename)

		key = strings.ReplaceAll(key, " ", "-")

		param.PhotoFile.Filename = key

		photoUrl, err = r.supabase.Upload(&param.PhotoFile)
		if err != nil {
			return res, err
		}
	}

	inputParam := entity.ReportInputParam{
		Title:       param.Title,
		Description: param.Description,
		Category:    entity.Category(param.Category),
		Location:    param.Location,
		PhotoUrl:    null.StringFrom(photoUrl),
		TicketCode:  ticketCode,
	}

	err = r.report.Create(ctx, inputParam)
	if err != nil {
		return res, err
	}

	res.TicketCode = ticketCode

	return res, err
}

func (r *report) randString() (string, error) {
	var b strings.Builder
	const alphabet = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"

	randLenght := 6
	b.Grow(randLenght)
	max := big.NewInt(int64(len(alphabet)))
	for i := 0; i < randLenght; i++ {
		r, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		b.WriteByte(alphabet[r.Int64()])
	}
	return b.String(), nil
}

func (r *report) GenerateTicketCode(t time.Time, loc *time.Location) (string, error) {
	if loc != nil {
		t = t.In(loc)
	}
	date := t.Format("20060102")
	rs, err := r.randString()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s-%s", date, rs), nil
}

func (r *report) GetAllReports(ctx context.Context, param dto.ReportParam) ([]dto.AllReports, error) {
	var reportsDto []dto.AllReports

	var from, to time.Time

	if !param.CreatedAt.IsZero() {
		from = time.Date(param.CreatedAt.Year(), param.CreatedAt.Month(), param.CreatedAt.Day(), 0, 0, 0, 0, wib)
		to = from.Add(24 * time.Hour)
	}

	reports, err := r.report.GetAll(ctx, entity.ReportParam{
		Title:           param.Title,
		Category:        param.Category,
		Status:          param.Status,
		CreatedAtHigher: from,
		CreatedAtLower:  to,
		PaginationParam: param.PaginationParam,
	})
	if err != nil {
		return reportsDto, err
	}

	for _, r := range reports {
		reportsDto = append(reportsDto, dto.AllReports{
			Id:         r.Id,
			TicketCode: r.TicketCode,
			Title:      r.Title,
			Category:   string(r.Category),
			Location:   r.Location,
			CreatedAt:  r.CreatedAt,
			Status:     string(r.Status),
		})
	}

	return reportsDto, err
}

func (r *report) GetReport(ctx context.Context, param dto.ReportParam) (dto.GetReport, error) {
	var reportRes dto.GetReport

	report, err := r.report.Get(ctx, entity.ReportParam{Id: param.Id})
	if err != nil {
		return reportRes, err
	}

	reportRes.Id = report.Id
	reportRes.Category = string(report.Category)
	reportRes.Title = report.Title
	reportRes.Description = report.Description
	reportRes.TicketCode = report.TicketCode
	reportRes.Location = report.Location
	reportRes.PhotoUrl = report.PhotoUrl.String
	reportRes.Status = string(report.Status)
	reportRes.StatusDesc = report.StatusDesc.String
	reportRes.StatusProofUrl = report.StatusProofUrl.String
	reportRes.CreatedAt = report.CreatedAt
	reportRes.UpdatedAt = report.UpdatedAt

	return reportRes, err
}

func (r *report) UpdateReport(ctx context.Context, param dto.UpdateParam) error {
	var proofPhotoUrl string
	var err error

	status := entity.Status(param.Status)

	switch status {
	case entity.INREVIEW:
		if strings.TrimSpace(param.StatusDesc) != "" || param.StatusProofFile.Size != 0 {
			return errors.NewWithCode(codes.CodeBadRequest, "in review didnt need any other properties")
		}
	case entity.REJECTED, entity.RESOLVED:
		if strings.TrimSpace(param.StatusDesc) == "" {
			return errors.NewWithCode(codes.CodeBadRequest, "status desc can't be empty")
		}
	}

	if param.StatusProofFile.Size != 0 && status != entity.INREVIEW {
		now := time.Now().In(wib)

		key := fmt.Sprintf("%s/%s/%d-%s", now.Format("02-01-2006"), status, now.UnixNano(), param.StatusProofFile.Filename)

		key = strings.ReplaceAll(key, " ", "-")

		param.StatusProofFile.Filename = key

		proofPhotoUrl, err = r.supabase.Upload(&param.StatusProofFile)
		if err != nil {
			return err
		}
	}

	updateParam := entity.UpdateReportParam{
		Status:         entity.Status(param.Status),
		StatusDesc:     null.StringFrom(param.StatusDesc),
		StatusProofUrl: null.StringFrom(proofPhotoUrl),
	}

	err = r.report.Update(ctx, updateParam, entity.ReportParam{Id: param.Id})
	if err != nil {
		return err
	}

	return nil
}
