package grpcServer

import (
	"context"
	"fmt"
	"management-backend/db"
	botpc "management-backend/grpc/bot"
	"management-backend/models"
	"management-backend/responses"
	"math/rand"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*gRPCServer) StartOperatorShift(ctx context.Context, req *botpc.OperatorRequest) (*botpc.OperatorShiftResponse, error) {
	// Cannot create after 22:00
	if time.Now().Hour() >= 22 {
		return nil, status.Errorf(
			codes.Unavailable,
			responses.ErrorNotToday,
		)
	}

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

	// Check if operator already on shift
	if len(operator.Shifts) > 0 && operator.Shifts[len(operator.Shifts)-1].EndDate.IsZero() {
		return nil, status.Errorf(
			codes.AlreadyExists,
			responses.ErrorAlreadyOnShift,
		)
	}

	// Check if operator have shift today
	if len(operator.Shifts) > 0 && operator.Shifts[len(operator.Shifts)-1].StartDate.Format("2006-01-02") == time.Now().Format("2006-01-02") {
		return nil, status.Errorf(
			codes.AlreadyExists,
			responses.ErrorAlreadyExists,
		)
	}

	// Create a new shift
	shift := models.Shift{
		StartDate:  time.Now(),
		EndDate:    time.Time{},
		LastNotify: time.Time{},
		NextNotify: time.Now().Add((time.Duration(rand.Intn(60-35) + 35)) * time.Minute),
	}

	if err := db.Model(&operator).Association("Shifts").Append(&shift); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	// Return operator
	return &botpc.OperatorShiftResponse{Shift: responses.CreateOperatorShiftGRPCResponse(&shift)}, nil
}
