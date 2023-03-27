package receiver

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

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
	config, err := readInConfig(*configPath)
	if err != nil {
		return fmt.Errorf("fail to read configuration: %w", err)
	}
	return startGrpcServer(config)
}

func readInConfig(path string) (*config.Config, error) {
	viper := viper.New()
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		return nil, fmt.Errorf("error happened while reading config file: %w", err)
	}
	config := &config.Config{}
	_ = viper.UnmarshalKey("receiver", config)
	return config, nil
}

func startGrpcServer(config *config.Config) error {
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(config.Port))
	if err != nil {
		log.Fatalf("Fail to listen Grpc Port: %v\n", err)
		return err
	}

	server := grpc.NewServer()
	// TraceId 接口
	traceIdServer := traceid.NewTraceIdServer(config.TraceIdCacheTime)
	model.RegisterTraceIdServiceServer(server, traceIdServer)

	traceIdServer.Start()
	log.Printf("Start Grpc Server: %d", config.Port)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Fail to start server: %v", err)
		return err
	}
	return nil
}
