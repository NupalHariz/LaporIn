package usecase

import (
	"github.com/nupalHariz/LaporIn/src/business/domain"
	"github.com/nupalHariz/LaporIn/src/business/service/supabase"
	"github.com/nupalHariz/LaporIn/src/business/usecase/report"
	"github.com/nupalHariz/LaporIn/src/business/usecase/user"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
	"github.com/reyhanmichiels/go-pkg/v2/hash"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/parser"
)

type Usecases struct {
	User   user.Interface
	Report report.Interface
}

type InitParam struct {
	Dom      *domain.Domains
	Json     parser.JSONInterface
	Log      log.Interface
	Hash     hash.Interface
	Auth     auth.Interface
	Supabase supabase.Interface
}

func Init(param InitParam) *Usecases {
	return &Usecases{
		User:   user.Init(user.InitParam{UserDomain: param.Dom.User, Auth: param.Auth, Hash: param.Hash}),
		Report: report.Init(report.InitParam{Report: param.Dom.Report, Supabase: param.Supabase}),
	}
}
