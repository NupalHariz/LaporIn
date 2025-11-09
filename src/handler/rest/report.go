package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nupalHariz/LaporIn/src/business/dto"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
)

func (r *rest) InputReport(ctx *gin.Context) {
	var param dto.InputReport

	err := r.Bind(ctx, &param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	res, err := r.uc.Report.InputReport(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeCreated, res, nil)
}

func (r *rest) GetAllReports(ctx *gin.Context) {
	res, err := r.uc.Report.GetAllReports(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeAccepted, res, nil)
}

func (r *rest) GetReport(ctx *gin.Context) {
	var param dto.ReportParam

	if err := r.BindUri(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	res, err := r.uc.Report.GetReport(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeAccepted, res, nil)
}

func (r *rest) UpdateReport(ctx *gin.Context) {
	var param dto.UpdateParam

	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.BindUri(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	err := r.uc.Report.UpdateReport(ctx, param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeAccepted, nil, nil)

}
