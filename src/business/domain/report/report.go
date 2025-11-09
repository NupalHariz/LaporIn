package report

import (
	"context"

	"github.com/nupalHariz/LaporIn/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

type Interface interface {
	Create(ctx context.Context, inputParam entity.ReportInputParam) error
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
