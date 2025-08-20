package di

import "assessment/domain/service"

var (
	planningService *service.PlanningService
)

func PlanningService() *service.PlanningService {
	if planningService != nil {
		return planningService
	}

	planningService = service.NewPlanningService(DbDatasource(), DbDatasource())

	return planningService
}
