package grpcServer

import (
	"context"
	"management-backend/db"
	botpc "management-backend/grpc/bot"
	"management-backend/models"
	"management-backend/responses"
	"math/rand"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*gRPCServer) WriteOperatorCaptcha(ctx context.Context, req *botpc.WriteOperatorCaptchaRequest) (*botpc.Empty, error) {
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
	if len(operator.Shifts) == 0 || !operator.Shifts[len(operator.Shifts)-1].EndDate.IsZero() || operator.Shifts[len(operator.Shifts)-1].LastNotify.IsZero() {
		return nil, status.Errorf(
			codes.NotFound,
			responses.ErrorNoShift,
		)
	}
	lastShift := operator.Shifts[len(operator.Shifts)-1]

	// Check is answer correct
	if strings.ToLower(lastShift.CaptchaAnswer) != strings.ToLower(req.Captcha) {
		return nil, status.Errorf(
			codes.InvalidArgument,
			responses.ErrorWrongAnswer,
		)
	}

	db.Model(&lastShift).Select("LastNotify", "CaptchaAnswer", "NextNotify").Updates(&models.Shift{
		LastNotify:    time.Time{},
		CaptchaAnswer: "",
		NextNotify:    time.Now().Add((time.Duration(rand.Intn(60-35) + 35)) * time.Minute),
	})

	return &botpc.Empty{}, nil
}
