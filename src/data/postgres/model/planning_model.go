package model

import (
	"time"

	"assessment/domain/entity"
)

type Planning struct {
	Id        int64
	Equipment int64
	Quantity  int64
	StartAt   time.Time
	EndAt     time.Time
}

func (m *Planning) ConvertEntityToModel(e *entity.Planning) {
	m.Id = e.Id
	m.Equipment = e.Equipment
	m.Quantity = e.Quantity
	m.StartAt = e.StartAt
	m.EndAt = e.EndAt
}

func (m *Planning) ConvertModelToEntity() *entity.Planning {
	return &entity.Planning{
		Id:        m.Id,
		Equipment: m.Equipment,
		Quantity:  m.Quantity,
		StartAt:   m.StartAt,
		EndAt:     m.EndAt,
	}
}
