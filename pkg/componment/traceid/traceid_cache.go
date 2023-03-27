package traceid

import (
	"sync"
	"time"

	"github.com/KindlingProject/camera-receiver/pkg/model"
)

type traceIdCache struct {
	traceIds sync.Map // <traceId, time>
	timeout  int64
}

func NewTraceIdCache(timeout int) *traceIdCache {
	return &traceIdCache{
		timeout: int64(timeout) * int64(time.Second),
	}
}

func (cache *traceIdCache) cacheTraceId(traceId string) {
	// 保留第一条TraceId 时间，减少TraceId数据被接口重复获取.
	cache.traceIds.LoadOrStore(traceId, time.Now().UnixNano())
}

func (cache *traceIdCache) cleanCache() {
	timer := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer.C:
			now := time.Now().UnixNano()
			cache.traceIds.Range(func(k, v any) bool {
				time := v.(int64)
				if now-time > cache.timeout {
					// 每秒钟清除已失效TraceId记录
					cache.traceIds.Delete(k)
				}
				return true
			})
		}
	}
}

func (cache *traceIdCache) getTraceIds(ignoreTraces map[string]bool, startTime int64) *model.TraceIds {
	traceIds := make([]string, 0)
	var lastTime int64
	cache.traceIds.Range(func(k, v any) bool {
		time := v.(int64)
		traceId := k.(string)
		if time > startTime {
			// 获取最新未获取TraceId记录
			if _, exist := ignoreTraces[traceId]; !exist {
				// 过滤自己发送的数据
				traceIds = append(traceIds, traceId)
			}
		}
		if time > lastTime {
			// 记录最大时间
			lastTime = time
		}
		return true
	})
	return &model.TraceIds{
		QueryTime: lastTime,
		TraceIds:  traceIds,
	}
}
