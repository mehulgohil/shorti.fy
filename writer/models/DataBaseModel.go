package models

import "time"

type URLTable struct {
	HashKey        string
	LongURL        string
	CreatedAt      time.Time
	ExpirationDate time.Time
	HitCount       int
	CreatedBy      string
}
