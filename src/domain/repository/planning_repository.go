package repository

import (
	"assessment/domain/entity"
	"context"
)

type PlanningRepository interface {
	GetPlannings(ctx context.Context) ([]entity.Planning, error)
}
