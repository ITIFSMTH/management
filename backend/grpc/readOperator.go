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

func (*gRPCServer) ReadOperator(ctx context.Context, req *botpc.OperatorRequest) (*botpc.OperatorResponse, error) {
	// Get DB
	db := db.GetDB()

	// Get operator
	var operator models.Operator
	if res := db.Model(&operator).Where(&models.Operator{
		Telegram: req.GetUsername(),
	}).Preload("Worker").
		Preload("Worker.Role").First(&operator); res.Error != nil || res.RowsAffected == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			responses.ErrorNotExists,
		)
	}

	// Return operator
	return &botpc.OperatorResponse{Operator: responses.CreateOperatorGRPCResponse(&operator)}, nil
}
