package responses

import (
	botpc "management-backend/grpc/bot"
	"management-backend/models"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Create a GRPC telegram bot key response
func CreateTelegramBotKeyGRPCResponse(settings *models.Setting) *botpc.TelegramBotKeyResponse {
	settingsResponse := CreateSettingsResponse(settings)

	return &botpc.TelegramBotKeyResponse{
		Key: settingsResponse.TelegramBotKey,
	}
}

// Create a GRPC worker role response
func CreateWorkerRoleGRPCResponse(workerRole *models.WorkerRole) *botpc.WorkerRole {
	workerRoleResponse := CreateWorkerRoleResponse(workerRole)

	return &botpc.WorkerRole{
		Id:   uint32(workerRoleResponse.ID),
		Role: workerRoleResponse.Role,
	}
}

// Create a GRPC worker response
func CreateWorkerGRPCResponse(worker *models.Worker) *botpc.Worker {
	workerResponse := CreateWorkerResponse(worker)

	return &botpc.Worker{
		Id:    uint32(workerResponse.ID),
		Login: workerResponse.Login,
		Role:  CreateWorkerRoleGRPCResponse(&worker.Role),
	}
}

// Create a GRPC operator response
func CreateOperatorGRPCResponse(operator *models.Operator) *botpc.Operator {
	operatorResponse := CreateOperatorResponse(operator)

	return &botpc.Operator{
		Id:         uint32(operatorResponse.ID),
		Worker:     CreateWorkerGRPCResponse(&operator.Worker),
		Telegram:   operatorResponse.Telegram,
		TelegramId: operator.TelegramID,
		OnShift:    operatorResponse.OnShift,
		OnTimeout:  operatorResponse.OnTimeout,
	}
}

// Create a GRPC operator shift response
func CreateOperatorShiftGRPCResponse(shift *models.Shift) *botpc.OperatorShift {
	operatorShiftResponse := CreateOperatorShiftResponse(shift)

	return &botpc.OperatorShift{
		StartDate: timestamppb.New(operatorShiftResponse.StartDate),
		EndDate:   timestamppb.New(operatorShiftResponse.EndDate),
		Delays:    uint32(shift.Delays),
	}
}

// Create a GRPC operator timeout response
func CreateOperatorTimeoutGRPCResponse(timeout *models.Timeout) *botpc.OperatorTimeout {
	operatorTimeoutResponse := CreateOperatorTimeoutResponse(timeout)

	return &botpc.OperatorTimeout{
		StartDate: timestamppb.New(operatorTimeoutResponse.StartDate),
		EndDate:   timestamppb.New(operatorTimeoutResponse.EndDate),
	}
}

// Create a GRPC operator timeouts response
func CreateOperatorTimeoutsGRPCResponse(timeouts *[]models.Timeout) []*botpc.OperatorTimeout {
	operatorTimeouts := []*botpc.OperatorTimeout{}

	for _, timeout := range *timeouts {
		operatorTimeoutResponse := CreateOperatorTimeoutResponse(&timeout)

		operatorTimeouts = append(operatorTimeouts, &botpc.OperatorTimeout{
			StartDate: timestamppb.New(operatorTimeoutResponse.StartDate),
			EndDate:   timestamppb.New(operatorTimeoutResponse.EndDate),
		})
	}

	return operatorTimeouts
}
