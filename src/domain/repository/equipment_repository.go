package repository

import (
	"assessment/domain/entity"
	"context"
)

type EquipmentRepository interface {
	GetEquipments(ctx context.Context) ([]entity.Equipment, error)
	GetEquipment(ctx context.Context, equipment int64) (*entity.Equipment, error)
}
