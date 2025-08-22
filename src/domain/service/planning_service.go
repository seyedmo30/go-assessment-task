package service

import (
	"assessment/domain/entity"
	"assessment/domain/repository"
	"assessment/dto"
	"context"
	"fmt"
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

// IsAvailable simplified and more readable.
// It uses the repository methods (same as your code) and the helpers above.
func (s *PlanningService) IsAvailable(equipment, quantity int64, startAt, endAt time.Time) (bool, error) {
	ctx := context.Background()

	// load equipment (stock)
	eq, err := s.equipmentRepository.GetEquipment(ctx, equipment)
	if err != nil {
		return false, err
	}
	if eq == nil {
		return false, fmt.Errorf("equipment %d not found", equipment)
	}

	// load only the plannings that overlap the requested window
	req := dto.GetPlanningsRequest{
		Equipment: equipment,
		StartAt:   startAt,
		EndAt:     endAt,
	}
	plannings, err := s.planningRepository.GetPlannings(ctx, req)
	if err != nil {
		return false, err
	}

	// build events and compute maximum concurrent usage
	events := buildEvents(plannings)
	maxConcurrent := maxConcurrentUsage(events)

	// final availability check
	// If at any moment the already-planned items + requested quantity exceed stock -> not available
	isAvailable := (maxConcurrent + quantity) <= eq.Stock
	return isAvailable, nil
}

// assume PlanningService has fields:
// planningRepository with method GetPlanningsBetween(ctx, startAt, endAt)
// equipmentRepository with method GetEquipments(ctx)
func (s *PlanningService) GetShortages(startAt, endAt time.Time) (map[int64]int64, error) {
	ctx := context.Background()

	// load all equipments
	equipments, err := s.equipmentRepository.GetEquipments(ctx)
	if err != nil {
		return nil, err
	}
	if len(equipments) == 0 {
		// no equipments -> return empty map
		return map[int64]int64{}, nil
	}

	getPlanningsBetweenRequestDTO:=dto.GetPlanningsBetweenRequest{
		StartAt: startAt,
		EndAt: endAt,
	}
	// load all plannings that overlap the window (single query for efficiency)
	allPlannings, err := s.planningRepository.GetPlanningsBetween(ctx,getPlanningsBetweenRequestDTO )
	if err != nil {
		return nil, err
	}

	// group plannings by equipment id
	grouped := make(map[int64][]entity.Planning)
	for _, p := range allPlannings {
		grouped[p.Equipment] = append(grouped[p.Equipment], p)
	}

	// for each equipment compute maxConcurrent and shortage
	result := make(map[int64]int64)
	for _, eq := range equipments {
		plans := grouped[eq.Id]
		if len(plans) == 0 {
			// nothing planned in range -> no shortage; skip (per requirement)
			continue
		}
		events := buildEvents(plans)
		maxConcurrent := maxConcurrentUsage(events)

		shortage := eq.Stock - maxConcurrent
		if shortage < 0 {
			// keep negative value as requested in spec
			result[eq.Id] = shortage
		}
	}

	return result, nil
}
