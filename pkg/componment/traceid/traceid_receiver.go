package traceid

import (
	"context"

	"github.com/KindlingProject/camera-receiver/pkg/model"
)

type TraceIdServer struct {
	model.UnimplementedTraceIdServiceServer
	cache *traceIdCache
}

func NewTraceIdServer(cacheTime int) *TraceIdServer {
	return &TraceIdServer{
		cache: NewTraceIdCache(cacheTime),
	}
}

func (server *TraceIdServer) Start() {
	go server.cache.cleanCache()
}

func (server *TraceIdServer) SendTraceIds(ctx context.Context, traceIds *model.TraceIds) (*model.TraceIds, error) {
	ignoreTraceIds := make(map[string]bool, 0)
	if traceIds != nil && len(traceIds.TraceIds) > 0 {
		for _, traceId := range traceIds.TraceIds {
			server.cache.cacheTraceId(traceId)
			// 标识查询时不需要该TraceId数据
			ignoreTraceIds[traceId] = true
		}
	}

	// 查询缓存数据
	return server.cache.getTraceIds(ignoreTraceIds, traceIds.QueryTime), nil
}
