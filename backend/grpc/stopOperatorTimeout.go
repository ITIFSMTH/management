package grpcServer

import (
	"context"
	"management-backend/db"
	botpc "management-backend/grpc/bot"
	"management-backend/models"
	"management-backend/responses"
	"math/rand"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*gRPCServer) StopOperatorTimeout(ctx context.Context, req *botpc.OperatorRequest) (*botpc.OperatorTimeoutResponse, error) {
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

	// Check is operator have opened shift
	if len(operator.Shifts) == 0 || (len(operator.Shifts) > 0 && !operator.Shifts[len(operator.Shifts)-1].EndDate.IsZero()) {
		return nil, status.Errorf(
			codes.NotFound,
			responses.ErrorNoShift,
		)
	}
	lastShift := operator.Shifts[len(operator.Shifts)-1]

	// Check is operator have opened timeout
	if len(lastShift.Timeouts) == 0 || (len(lastShift.Timeouts) > 0 && !lastShift.Timeouts[len(lastShift.Timeouts)-1].EndDate.IsZero()) {
		return nil, status.Errorf(
			codes.NotFound,
			responses.ErrorNoTimeout,
		)
	}
	lastTimeout := lastShift.Timeouts[len(lastShift.Timeouts)-1]

	// Create new time next notify
	db.Model(&lastShift).Updates(&models.Shift{
		NextNotify: time.Now().Add((time.Duration(rand.Intn(60-35) + 35)) * time.Minute),
	})

	// Close timeout
	db.Model(&lastTimeout).Updates(&models.Timeout{
		EndDate: time.Now(),
	})

	// Return operator
	return &botpc.OperatorTimeoutResponse{Timeout: responses.CreateOperatorTimeoutGRPCResponse(&lastTimeout)}, nil
}
