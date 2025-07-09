package service

import (
	"math"
    "gorm.io/gorm"
	"github.com/stilln0thing/WareManager/server/internal/model"
)

func ReassignAisles(db *gorm.DB, numAisles int) error {

	var skus []model.SKU
	if err := db.
		Order("pick_count DESC").
		Find(&skus).Error; err != nil {
			return err
	}
	total := len(skus)
	if total == 0{
		return nil
	}
	perAisle := int(math.Ceil(float64(total) / float64(numAisles)))
	return db.Transaction(func(tx *gorm.DB) error {
        for idx, sku := range skus {
           
            aisle := (idx / perAisle) + 1
            if aisle > numAisles {
                aisle = numAisles
            }

           
            if err := tx.
                Model(&model.SKUItem{}).
                Where("sk_uid = ?", sku.ID).
                Update("aisle_no", aisle).
                Error; err != nil {
                return err
            }
        }
        return nil
    })
}