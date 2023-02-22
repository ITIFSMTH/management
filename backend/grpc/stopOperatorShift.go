package grpcServer

import (
	"context"
	"management-backend/db"
	botpc "management-backend/grpc/bot"
	"management-backend/models"
	"management-backend/responses"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*gRPCServer) StopOperatorShift(ctx context.Context, req *botpc.OperatorRequest) (*botpc.OperatorShiftWithTimeoutsResponse, error) {
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

	// Check is operator have opened shift
	if len(operator.Shifts) == 0 || (len(operator.Shifts) > 0 && !operator.Shifts[len(operator.Shifts)-1].EndDate.IsZero()) {
		return nil, status.Errorf(
			codes.NotFound,
			responses.ErrorNoShift,
		)
	}
	lastShift := operator.Shifts[len(operator.Shifts)-1]

	// Get all shift timeouts
	db.Model(&lastShift).Preload("Timeouts").Find(&lastShift)

	// Close shift
	db.Model(&lastShift).Select("LastNotify", "NextNotify", "CaptchaAnswer", "EndDate").Updates(&models.Shift{
		EndDate:       time.Now(),
		LastNotify:    time.Time{},
		NextNotify:    time.Time{},
		CaptchaAnswer: "",
	})

	// Return operator
	return &botpc.OperatorShiftWithTimeoutsResponse{
		Shift:    responses.CreateOperatorShiftGRPCResponse(&lastShift),
		Timeouts: responses.CreateOperatorTimeoutsGRPCResponse(&lastShift.Timeouts),
	}, nil
}
