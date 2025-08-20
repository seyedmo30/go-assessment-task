package repository

import (
	"assessment/domain/entity"
	"context"
)

type EquipmentRepository interface {
	GetEquipments(ctx context.Context) ([]entity.Equipment, error)
}
