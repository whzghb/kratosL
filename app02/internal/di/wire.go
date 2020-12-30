// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"kratosL/app02/internal/dao"
	"kratosL/app02/internal/service"
	"kratosL/app02/internal/server/grpc"
	"kratosL/app02/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
