package grpcServer

import (
	"context"
	botpc "management-backend/grpc/bot"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type gRPCServer struct {
	botpc.BotServiceServer
}

func Run(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(serverLoggerInterceptor),
	)

	botpc.RegisterBotServiceServer(grpcServer, &gRPCServer{})
	grpcServer.Serve(listener)
}

// Logger unary interceptor function
func serverLoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Time of getting request
	start := time.Now()

	// Calls the handler
	h, err := handler(ctx, req)

	// Logging with grpclog (grpclog.LoggerV2)
	grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout).Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}
