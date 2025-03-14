package datamodel

import "time"

type Model struct {
	id          int       `json:"id"`
	description string    `json:"description"`
	status      string    `json:"status"`
	createdAt   time.Time `json:"createdAt"`
	updatedAt   time.Time `json:"updatedAt"`
}
