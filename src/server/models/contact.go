package models

import "gorm.io/gorm"

// TODO: Add foreign key logic to ContactInfo model
type ContactInfo struct {
	gorm.Model
	ID           uint `gorm:"primaryKey;serial"`
	OwnerID      uint
	AddressID    uint
	PhoneNumber1 string
	PhoneNumber2 string
	FaxNumber    string
}

type Address struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;serial"`
	Address1 string `gorm:"not null"`
	Address2 string
	City     string `gorm:"not null"`
	State    string `gorm:"not null"`
	ZipCode  string `gorm:"not null"`
}
