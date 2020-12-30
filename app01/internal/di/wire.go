// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"app01/internal/dao"
	"app01/internal/service"
	"app01/internal/server/grpc"
	"app01/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
