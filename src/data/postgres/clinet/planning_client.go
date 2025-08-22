package clinet

import (
	"assessment/data/postgres/model"
	"assessment/domain/entity"
	"assessment/dto"
	"context"
)

func (p *Postgres) GetPlannings(ctx context.Context, req dto.GetPlanningsRequest) ([]entity.Planning, error) {
	var models []model.Planning
	if err := p.db.WithContext(ctx).
		Where("equipment = ? AND start_at < ? AND end_at > ?", req.Equipment, req.EndAt, req.StartAt).
		Find(&models).Error; err != nil {
		return nil, err
	}
	res := make([]entity.Planning, 0, len(models))
	for _, m := range models {
		res = append(res, *m.ConvertModelToEntity())
	}
	return res, nil
}

func (p *Postgres) GetPlanningsBetween(ctx context.Context, req dto.GetPlanningsBetweenRequest) ([]entity.Planning, error) {
	var models []model.Planning
	if err := p.db.WithContext(ctx).
		Where("start_at < ? AND end_at > ?", req.EndAt, req.StartAt).
		Find(&models).Error; err != nil {
		return nil, err
	}

	res := make([]entity.Planning, 0, len(models))
	for _, m := range models {
		res = append(res, *m.ConvertModelToEntity())
	}
	return res, nil
}
