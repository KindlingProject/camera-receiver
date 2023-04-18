package p9x

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/KindlingProject/camera-receiver/pkg/model"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	prometheus_model "github.com/prometheus/common/model"
)

const (
	Instance    = "instance"
	ContentKey  = "content_key"
	ContainerId = "container_id"
)

type prometheusP9xCache struct {
	// Pod - Url - P9x
	p9xCache          map[string]*model.P9XResponse
	query             string
	prometheusAddress string
}

func NewPrometheusP9xCache(address string, p9xValue float32, outputName string, duration string) *prometheusP9xCache {
	query := fmt.Sprintf("histogram_quantile(%f, sum by (%s, %s, %s, le) (rate(%s[%s])))",
		p9xValue,
		Instance,
		ContainerId,
		ContentKey,
		outputName,
		duration,
	)
	return &prometheusP9xCache{
		p9xCache:          make(map[string]*model.P9XResponse),
		query:             query,
		prometheusAddress: address,
	}
}

func (cache *prometheusP9xCache) queryP9xByPromethues(interval int) {
	log.Printf("QueryP9x with Promethus in %d second.", interval)
	timer := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-timer.C:
			cache.updateP9xByPromethues()
		}
	}
}

func (cache *prometheusP9xCache) updateP9xByPromethues() {
	client, err := api.NewClient(api.Config{
		Address: cache.prometheusAddress,
	})
	if err != nil {
		log.Fatalf("NewClient for Promethues failed: %v", err)
		return
	}
	v1Api := v1.NewAPI(client)
	result, warnings, err := v1Api.Query(context.Background(), cache.query, time.Now())
	if err != nil {
		log.Fatalf("Request Prometheus P9x failed: %v", err)
		return
	}
	if len(warnings) > 0 {
		log.Printf("Request Prometheus Warning: %s", warnings)
	}

	p9xCache := make(map[string]*model.P9XResponse)
	var (
		cacheResponse *model.P9XResponse
		exist         bool
	)
	count := 0
	if vector, ok := result.(prometheus_model.Vector); ok {
		for _, sample := range vector {
			instance := string(sample.Metric[Instance])
			if cacheResponse, exist = p9xCache[instance]; !exist {
				cacheResponse = &model.P9XResponse{
					Datas: make([]*model.P9XData, 0),
				}
				p9xCache[instance] = cacheResponse
			}
			if float64(sample.Value) > 0 {
				count++
				cacheResponse.Datas = append(cacheResponse.Datas, &model.P9XData{
					Url:         string(sample.Metric[ContentKey]),
					ContainerId: string(sample.Metric[ContainerId]),
					Value:       float64(sample.Value),
				})
			}
		}
	}
	log.Printf("Receive [P9x] Count: %d, Group: %d\n", count, len(p9xCache))
	if count > 0 {
		cache.p9xCache = p9xCache
	}
}

func (cache *prometheusP9xCache) queryP9x(ip string) *model.P9XResponse {
	localCache := cache.p9xCache
	for key, value := range localCache {
		if strings.HasPrefix(key, ip) {
			return value
		}
	}
	return &model.P9XResponse{
		Datas: make([]*model.P9XData, 0),
	}
}
