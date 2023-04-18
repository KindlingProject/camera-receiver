package p9x

import (
	"context"

	"github.com/KindlingProject/camera-receiver/pkg/config"
	"github.com/KindlingProject/camera-receiver/pkg/model"
)

type P9XServer struct {
	model.UnimplementedP9XServiceServer
	cache         *prometheusP9xCache
	queryInterval int
}

func NewP9XServer(cfg *config.PrometheusConfig) *P9XServer {
	return &P9XServer{
		cache: NewPrometheusP9xCache(
			cfg.Address,
			cfg.QueryP9xVal,
			cfg.QueryP9xTable,
			cfg.QueryP9xDuration,
		),
		queryInterval: cfg.QueryInterval,
	}
}

func (server *P9XServer) Start() {
	go server.cache.queryP9xByPromethues(server.queryInterval)
}

func (server *P9XServer) QueryP9X(ctx context.Context, request *model.P9XRequest) (*model.P9XResponse, error) {
	return server.cache.queryP9x(request.Ip), nil
}
