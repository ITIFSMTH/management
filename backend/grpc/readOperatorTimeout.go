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

func (*gRPCServer) ReadOperatorTimeout(ctx context.Context, req *botpc.OperatorRequest) (*botpc.OperatorTimeoutResponse, error) {
	// Get DB
	db := db.GetDB()

	// Get operator
	var operator models.Operator
	if res := db.Model(&operator).Where(&models.Operator{
		Telegram: req.GetUsername(),
	}).Preload("Shifts").Preload("Shifts.Timeouts").First(&operator); res.Error != nil || res.RowsAffected == 0 {
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

	// Check is timeout exist
	if len(lastShift.Timeouts) == 0 || !lastShift.Timeouts[len(lastShift.Timeouts)-1].EndDate.IsZero() {
		return nil, status.Errorf(
			codes.NotFound,
			responses.ErrorNoShift,
		)
	}
	lastTimeout := lastShift.Timeouts[len(lastShift.Timeouts)-1]

	// Return operator
	return &botpc.OperatorTimeoutResponse{Timeout: responses.CreateOperatorTimeoutGRPCResponse(&lastTimeout)}, nil
}
