package dto

import "time"

type GetPlanningsRequest struct {
	Equipment int64
	StartAt   time.Time
	EndAt     time.Time
}

type GetPlanningsBetweenRequest struct {
	StartAt time.Time
	EndAt   time.Time
}
