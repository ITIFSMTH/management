package grpcServer

import (
	"context"
	"fmt"
	"management-backend/db"
	botpc "management-backend/grpc/bot"
	"management-backend/models"
	"management-backend/responses"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*gRPCServer) StartOperatorTimeout(ctx context.Context, req *botpc.OperatorRequest) (*botpc.OperatorTimeoutResponse, error) {
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

	// Check is current shift waiting a captcha asnwer
	if !lastShift.LastNotify.IsZero() {
		return nil, status.Errorf(
			codes.Aborted,
			responses.ErrorFirstProvideCaptcha,
		)
	}

	// Drop next notify time
	db.Model(&lastShift).Select("NextNotify").Updates(&models.Shift{
		NextNotify: time.Time{},
	})

	// Create a new timeout
	timeout := models.Timeout{
		StartDate: time.Now(),
		EndDate:   time.Time{},
	}

	if err := db.Model(&lastShift).Association("Timeouts").Append(&timeout); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	// Return operator
	return &botpc.OperatorTimeoutResponse{Timeout: responses.CreateOperatorTimeoutGRPCResponse(&timeout)}, nil
}
