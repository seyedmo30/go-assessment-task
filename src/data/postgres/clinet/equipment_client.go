package clinet

import (
	"assessment/data/postgres/model"
	"assessment/domain/entity"
	"context"
)

func (p *Postgres) GetEquipments(ctx context.Context) ([]entity.Equipment, error) {
	var models []model.Equipment
	if err := p.db.WithContext(ctx).
		Find(&models).Error; err != nil {
		return nil, err
	}

	var res []entity.Equipment
	for _, m := range models {

		res = append(res, *m.ConvertModelToEntity())
	}
	return res, nil
}

func (p *Postgres) GetEquipment(ctx context.Context, equipment int64) (*entity.Equipment, error) {
	// TODO: error handle inside of repo
	var model model.Equipment
	if err := p.db.WithContext(ctx).
		Where("id = ?", equipment).
		First(&model).Error; err != nil {
		return nil, err
	}
	return model.ConvertModelToEntity(), nil
}
