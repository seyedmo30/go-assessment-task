package repository

import (
	"assessment/domain/entity"
	"assessment/dto"
	"context"
)

type PlanningRepository interface {
	GetPlannings(ctx context.Context, req dto.GetPlanningsRequest) ([]entity.Planning, error)
	GetPlanningsBetween(ctx context.Context, req dto.GetPlanningsBetweenRequest) ([]entity.Planning, error)
}
