package model

import (
	"time"
)

// this has all the SKU's we have
type SKU struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"unique;not null"`
    Category  string    `gorm:"not null"`
    PickCount int       `gorm:"not null;default:0"`
    CreatedAt time.Time
    UpdatedAt time.Time

    Items []SKUItem `gorm:"foreignKey:SKUId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// this will map to one SKU and give us the aisle where the SKU is present and the quantity of the SKU
// for now intime will be time.NOW()
type SKUItem struct {
    ID        uint     	`gorm:"primaryKey"`
    SKUId     uint      		`gorm:"not null"`   
    InTime    time.Time 		`gorm:"not null"`
    OutTime  *time.Time                        
    AisleNo   int       	    `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
	Quantity  uint              `gorm:"not null;default:0"`
    SKU SKU `gorm:"references:ID"` // Optional belongs-to relationship
}
