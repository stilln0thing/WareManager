package model

import "gorm.io/gorm"

type SKU struct {
	ID 			uint 				`gorm:"primaryKey"`
	Name 		string				`gorm:"unique;not null"`
	Qty 		int					`gorm:"not null;default:0"`
	PickCount   int			 		`gorm:"not null;default:0"`
	Aisle 		int					`gorm:"not null;default:0"`
	gorm.Model
}