package report

import (
	"context"

	"github.com/nupalHariz/LaporIn/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

type Interface interface {
	Create(ctx context.Context, inputParam entity.ReportInputParam) error
	GetAll(ctx context.Context, param entity.ReportParam) ([]entity.Report, error)
	Get(ctx context.Context, param entity.ReportParam) (entity.Report, error)
	Update(ctx context.Context, updateBody entity.UpdateReportParam, param entity.ReportParam) error
}

type report struct {
	db  sql.Interface
	log log.Interface
}

type InitParam struct {
	Db  sql.Interface
	Log log.Interface
}

func Init(param InitParam) Interface {
	return &report{
		db:  param.Db,
		log: param.Log,
	}
}

func (r *report) Create(ctx context.Context, inputParam entity.ReportInputParam) error {
	err := r.createSQL(ctx, inputParam)

	return err
}

func (r *report) GetAll(ctx context.Context, param entity.ReportParam) ([]entity.Report, error) {
	reports, err := r.getAllSQL(ctx, param)

	return reports, err
}

func (r *report) Get(ctx context.Context, param entity.ReportParam) (entity.Report, error) {
	report, err := r.getSQL(ctx, param)

	return report, err
}

func (r *report) Update(ctx context.Context, updateBody entity.UpdateReportParam, param entity.ReportParam) error {
	err := r.updateSQL(ctx, updateBody, param)
	if err != nil {
		return err
	}

	return err
}
