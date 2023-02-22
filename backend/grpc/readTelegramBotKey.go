package grpcServer

import (
	"context"
	"management-backend/db"
	botpc "management-backend/grpc/bot"
	"management-backend/models"
	"management-backend/responses"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*gRPCServer) ReadTelegramBotKey(ctx context.Context, req *botpc.Empty) (*botpc.TelegramBotKeyResponse, error) {
	// Get DB
	db := db.GetDB()

	// Get settings
	var settings models.Setting
	if res := db.Model(&settings).First(&settings); res.Error != nil || res.RowsAffected == 0 {
		return nil, status.Errorf(
			codes.Internal,
			responses.ErrorServer,
		)
	}

	// Return token
	return responses.CreateTelegramBotKeyGRPCResponse(&settings), nil
}
