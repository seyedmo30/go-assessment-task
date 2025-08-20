package service

import (
	"assessment/domain/repository"
	"errors"
	"time"
)

type PlanningService struct {
	equipmentRepository repository.EquipmentRepository
	planningRepository  repository.PlanningRepository
}

func NewPlanningService(equipmentRepository repository.EquipmentRepository, planningRepository repository.PlanningRepository) *PlanningService {
	return &PlanningService{
		equipmentRepository: equipmentRepository,
		planningRepository:  planningRepository,
	}
}

func (s *PlanningService) IsAvailable(equipment, quantity int64, startAt, endAt time.Time) (bool, error) {
	// TODO: implement
	return false, errors.New("not implemented yet")
}

func (s *PlanningService) GetShortages(startAt, endAt time.Time) (map[int64]int64, error) {
	// TODO: implement
	// This method should return a list of shortages based on the planning
	return nil, errors.New("not implemented yet")
}
