package model

import "assessment/domain/entity"

type Equipment struct {
	Id    int64
	Name  string
	Stock int64
}

func (m *Equipment) ConvertEntityToModel(e *entity.Equipment) {
	m.Id = e.Id
	m.Name = e.Name
}

func (m *Equipment) ConvertModelToEntity() *entity.Equipment {
	return &entity.Equipment{
		Id:    m.Id,
		Name:  m.Name,
		Stock: m.Stock,
	}
}
