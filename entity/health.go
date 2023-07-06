package entity

import "time"

type Health struct {
	Sync *time.Time `json:"sync"`
}
