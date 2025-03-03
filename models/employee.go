package models

import (
	"time"

	"gorm.io/gorm"
)

type M_Employee struct {
	gorm.Model
	M_Employee_ID     uint           `gorm:"primaryKey"`
	EmployeeID        string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	FirstName         string         `gorm:"type:varchar(100);not null"`
	LastName          string         `gorm:"type:varchar(100);not null"`
	Email             string         `gorm:"type:uniqueIndex;not null"`
	Phone             string         `gorm:"type:varchar(20)"`
	Address           string         `gorm:"type:varchar(255)"`
	NationalID        string         `gorm:"type:uniqueIndex;not null"`
	TaxID             string         `gorm:"type:uniqueIndex"`
	ProfilePicURL     string         `gorm:"type:varchar(255)"`
	PayStatus         string         `gorm:"type:varchar(3);not null"`
	PayStatus_Key     uint           `gorm:"not null"`
	BirthName         string         `gorm:"type:varchar(255)"`
	DateOfBirth       time.Time      `gorm:"not null"`
	Gender            string         `gorm:"type:varchar(3);not null"`
	Gender_Key        uint           `gorm:"not null"`
	GenderRef         AD_Reference   `gorm:"foreignKey=Gender_Key;references=AD_Reference_ID"`
	Nationality       string         `gorm:"type:varchar(3);not null"`
	Nationality_Key   uint           `gorm:"not null"`
	NationalityRef    AD_Reference   `gorm:"foreignKey=Nationality_Key;references=AD_Reference_ID"`
	BloodType         string         `gorm:"type:varchar(3);not null"`
	BloodType_Key     uint           `gorm:"not null"`
	BloodTypeRef      AD_Reference   `gorm:"foreignKey=BloodType_Key;references=AD_Reference_ID"`
	DatePayStart      time.Time      `gorm:"not null"`
	DatePayEnd        time.Time      `gorm:"null"`
	Department        string         `gorm:"type:varchar(3);not null"`
	Department_Key    uint           `gorm:"not null"`
	DepartmentRef     AD_Reference   `gorm:"foreignKey=Department_Key;references=AD_Reference_ID"`
	MaritalStatus     string         `gorm:"type:varchar(3);not null"`
	MaritalStatus_Key uint           `gorm:"not null"`
	MaritalStatusRef  AD_Reference   `gorm:"foreignKey=MaritalStatus_Key;references=AD_Reference_ID"`
	Religion          string         `gorm:"type:varchar(3);not null"`
	Religion_Key      uint           `gorm:"not null"`
	ReligionRef       AD_Reference   `gorm:"foreignKey=Religion_Key;references=AD_Reference_ID"`
	PayrollID         string         `gorm:"type:varchar(100)"`
	HiredAt           time.Time      `gorm:"not null"`
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
