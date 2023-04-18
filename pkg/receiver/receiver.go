package receiver

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/KindlingProject/camera-receiver/pkg/componment/p9x"
	"github.com/KindlingProject/camera-receiver/pkg/componment/traceid"
	"github.com/KindlingProject/camera-receiver/pkg/config"
	"github.com/KindlingProject/camera-receiver/pkg/model"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func Run(ctx context.Context) error {
	// Initialize flags
	configPath := flag.String("config", "receiver-config.yml", "Configuration file")
	flag.Parse()
	receiverCfg, prometheusCfg, err := readInConfig(*configPath)
	if err != nil {
		return fmt.Errorf("fail to read configuration: %w", err)
	}
	return startGrpcServer(receiverCfg, prometheusCfg)
}

func readInConfig(path string) (*config.ReceiverConfig, *config.PrometheusConfig, error) {
	viper := viper.New()
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		return nil, nil, fmt.Errorf("error happened while reading config file: %w", err)
	}
	receiverCfg := &config.ReceiverConfig{}
	_ = viper.UnmarshalKey("receiver", receiverCfg)

	prometheusCfg := &config.PrometheusConfig{}
	_ = viper.UnmarshalKey("promethues", prometheusCfg)
	return receiverCfg, prometheusCfg, nil
}

func startGrpcServer(receiverCfg *config.ReceiverConfig, prometheusCfg *config.PrometheusConfig) error {
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(receiverCfg.Port))
	if err != nil {
		log.Fatalf("Fail to listen Grpc Port: %v\n", err)
		return err
	}

	server := grpc.NewServer()
	// TraceId 接口
	log.Printf("Start TraceId Cache Time: %d", receiverCfg.TraceIdCacheTime)
	traceIdServer := traceid.NewTraceIdServer(receiverCfg.TraceIdCacheTime)
	model.RegisterTraceIdServiceServer(server, traceIdServer)
	traceIdServer.Start()

	log.Printf("Start P9x Address: %s", prometheusCfg.Address)
	p9xServer := p9x.NewP9XServer(prometheusCfg)
	model.RegisterP9XServiceServer(server, p9xServer)
	p9xServer.Start()

	log.Printf("Start Grpc Server: %d", receiverCfg.Port)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Fail to start server: %v", err)
		return err
	}
	return nil
}
