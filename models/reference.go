package models

import (
	"time"

	"gorm.io/gorm"
)

type AD_Reference struct {
	AD_Reference_ID uint           `gorm:"primaryKey"`
	ReferenceName   string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	ReferenceValue  string         `gorm:"type:varchar(100);not null"`
	ReferenceDesc   string         `gorm:"type:varchar(255);not null"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
