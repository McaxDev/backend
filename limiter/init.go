package limiter

import (
	"time"

	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	Limiter map[string][]LimitRule
)

type LimitRule struct {
	Count    uint
	Duration time.Duration
}
