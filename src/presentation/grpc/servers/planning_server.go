package servers

import (
	"assessment/domain/service"
	"assessment/presentation/grpc/pbs"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type PlanningServer struct {
	service *service.PlanningService
}

func NewPlanningServer(service *service.PlanningService) *PlanningServer {
	return &PlanningServer{
		service: service,
	}
}

func (s *PlanningServer) IsAvailable(ctx context.Context, request *pbs.AvailabilityInquiryRequest) (*pbs.AvailabilityInquiryResponse, error) {
	if request.GetStartAt() == "" {
		return nil, status.New(codes.InvalidArgument, "empty start_at").Err()
	}
	if request.GetEndAt() == "" {
		return nil, status.New(codes.InvalidArgument, "empty end_at").Err()
	}

	start, err := time.Parse("2006-01-02 15:04:05", request.GetStartAt())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid start_at format").Err()
	}

	end, err := time.Parse("2006-01-02 15:04:05", request.GetEndAt())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid end_at format").Err()
	}

	available, err := s.service.IsAvailable(request.GetEquipment(), request.GetQuantity(), start, end)
	if err != nil {
		return nil, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	return &pbs.AvailabilityInquiryResponse{
		IsAvailable: available,
	}, nil
}

func (s *PlanningServer) GetShortages(ctx context.Context, request *pbs.GetShortagesInquiryRequest) (*pbs.GetShortagesInquiryResponse, error) {
	if request.GetStartAt() == "" {
		return nil, status.New(codes.InvalidArgument, "empty start_at").Err()
	}
	if request.GetEndAt() == "" {
		return nil, status.New(codes.InvalidArgument, "empty end_at").Err()
	}

	start, err := time.Parse("2006-01-02 15:04:05", request.GetStartAt())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid start_at format").Err()
	}

	end, err := time.Parse("2006-01-02 15:04:05", request.GetEndAt())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid end_at format").Err()
	}

	shortages, err := s.service.GetShortages(start, end)
	if err != nil {
		return nil, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	return &pbs.GetShortagesInquiryResponse{
		Shortages: shortages,
	}, nil
}
