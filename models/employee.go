package models

import (
	"time"

	"gorm.io/gorm"
)

type M_Employee struct {
	M_Employee_ID uint           `gorm:"primaryKey"`
	EmployeeID    string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	FirstName     string         `gorm:"type:varchar(100);not null"`
	LastName      string         `gorm:"type:varchar(100);not null"`
	Email         string         `gorm:"type:uniqueIndex;not null"`
	Phone         string         `gorm:"type:varchar(20)"`
	Address       string         `gorm:"type:varchar(255)"`
	NationalID    string         `gorm:"type:uniqueIndex;not null"`
	TaxID         string         `gorm:"type:uniqueIndex"`
	ProfilePicURL string         `gorm:"type:varchar(255)"`
	PayStatus     string         `gorm:"type:varchar(3);not null"`
	BirthName     string         `gorm:"type:varchar(255)"`
	DateOfBirth   time.Time      `gorm:"not null"`
	Gender        string         `gorm:"type:varchar(3);not null"`
	Nationality   string         `gorm:"type:varchar(3);not null"`
	BloodType     string         `gorm:"type:varchar(3);not null"`
	DatePayStart  time.Time      `gorm:"not null"`
	DatePayEnd    time.Time      `gorm:"null"`
	Department    string         `gorm:"type:varchar(3);not null"`
	PayrollID     string         `gorm:"type:varchar(100)"`
	HiredAt       time.Time      `gorm:"not null"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
