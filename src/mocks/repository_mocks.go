package mocks

import (
	"context"

	"assessment/domain/entity"
	"assessment/dto"

	"github.com/stretchr/testify/mock"
)

type MockEquipmentRepo struct{ mock.Mock }

func (m *MockEquipmentRepo) GetEquipment(ctx context.Context, id int64) (*entity.Equipment, error) {
	args := m.Called(ctx, id)
	if e := args.Get(0); e != nil {
		return e.(*entity.Equipment), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockEquipmentRepo) GetEquipments(ctx context.Context) ([]entity.Equipment, error) {
	args := m.Called(ctx)
	if v := args.Get(0); v != nil {
		return v.([]entity.Equipment), args.Error(1)
	}
	return nil, args.Error(1)
}

type MockPlanningRepo struct{ mock.Mock }

func (m *MockPlanningRepo) GetPlannings(ctx context.Context, req dto.GetPlanningsRequest) ([]entity.Planning, error) {
	args := m.Called(ctx, req)
	if v := args.Get(0); v != nil {
		return v.([]entity.Planning), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockPlanningRepo) GetPlanningsBetween(ctx context.Context, req dto.GetPlanningsBetweenRequest) ([]entity.Planning, error) {
	args := m.Called(ctx, req)
	if v := args.Get(0); v != nil {
		return v.([]entity.Planning), args.Error(1)
	}
	return nil, args.Error(1)
}
