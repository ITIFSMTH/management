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

func (*gRPCServer) WriteOperatorTelegramID(ctx context.Context, req *botpc.WriteTelegramIDRequest) (*botpc.Empty, error) {
	// Get DB
	db := db.GetDB()

	// Update telegram ID
	if res := db.Model(&models.Operator{}).Where("(telegram_id IS NULL OR telegram_id = 0) AND telegram = ?", req.Username).Update("telegram_id", req.TelegramId); res.Error != nil || res.RowsAffected == 0 {
		return nil, status.Errorf(
			codes.AlreadyExists,
			responses.ErrorAlreadyExists,
		)
	}

	return &botpc.Empty{}, nil
}
