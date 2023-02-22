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

func (*gRPCServer) ReadOperatorShift(ctx context.Context, req *botpc.OperatorRequest) (*botpc.OperatorShiftResponse, error) {
	// Get DB
	db := db.GetDB()

	// Get operator
	var operator models.Operator
	if res := db.Model(&operator).Where(&models.Operator{
		Telegram: req.GetUsername(),
	}).Preload("Shifts").First(&operator); res.Error != nil || res.RowsAffected == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			responses.ErrorNotExists,
		)
	}

	// Check is shift exist
	if len(operator.Shifts) == 0 || !operator.Shifts[len(operator.Shifts)-1].EndDate.IsZero() {
		return nil, status.Errorf(
			codes.NotFound,
			responses.ErrorNoShift,
		)
	}
	lastShift := operator.Shifts[len(operator.Shifts)-1]

	// Return operator
	return &botpc.OperatorShiftResponse{Shift: responses.CreateOperatorShiftGRPCResponse(&lastShift)}, nil
}
