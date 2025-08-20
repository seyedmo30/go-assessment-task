package entity

import "time"

type Planning struct {
	Id        int64
	Equipment int64
	Quantity  int64
	StartAt   time.Time
	EndAt     time.Time
}
