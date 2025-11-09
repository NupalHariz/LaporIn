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
	"github.com/reyhanmichiels/go-pkg/v2/null"
)

type Interface interface {
	InputReport(ctx context.Context, inputParam dto.InputReport) (dto.InputReportResponse, error)
	GetAllReports(ctx context.Context) ([]dto.AllReports, error)
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

func (r *report) InputReport(ctx context.Context, param dto.InputReport) (dto.InputReportResponse, error) {
	res := dto.InputReportResponse{}

	ticketCode, err := r.GenerateTicketCode(time.Now(), time.FixedZone("WIB", 7*3600))
	if err != nil {
		return res, err
	}

	var photoUrl string

	if param.PhotoFile.Size != 0 {
		fmt.Println("MASUK SINI")
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

func (r *report) GetAllReports(ctx context.Context) ([]dto.AllReports, error) {
	var reportsDto []dto.AllReports

	reports, err := r.report.GetAll(ctx)
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
		})
	}

	return reportsDto, err
}
