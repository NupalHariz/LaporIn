package report

import (
	"context"
	"strings"

	"github.com/nupalHariz/LaporIn/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

func (r *report) createSQL(ctx context.Context, inputParam entity.ReportInputParam) error {
	tx, err := r.db.Leader().BeginTx(ctx, "txReport", sql.TxOptions{})
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("iNewReport", insertReport, inputParam)
	if err != nil {
		if strings.Contains(err.Error(), entity.DuplicateEntryErrMessage) {
			return errors.NewWithCode(codes.CodeSQLUniqueConstraint, err.Error())
		}

		return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if rowCount < 1 {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no budget created")
	}

	if err := tx.Commit(); err != nil {
		return errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return err
}

func (r *report) getAllSQL(ctx context.Context) ([]entity.Report, error) {
	var reports []entity.Report

	var x []interface{}

	rows, err := r.db.Query(ctx, "rReportAll", getReport, x...)
	if err != nil {
		return reports, err
	}
	defer rows.Close()

	for rows.Next() {
		report := entity.Report{}

		err := rows.StructScan(&report)
		if err != nil {
			r.log.Error(ctx, codes.CodeSQLRowScan)
			continue
		}

		reports = append(reports, report)
	}

	return reports, nil
}
