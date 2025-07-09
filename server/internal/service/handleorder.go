package service

import (
	"github.com/stilln0thing/WareManager/server/internal/db" 
    "fmt"
    "strconv"
	"github.com/stilln0thing/WareManager/server/internal/model"

    "gorm.io/gorm"
)

func HandleOrder() error {
	cart, err := FetchCart()
	if err != nil {
		return err
	}
	
	return db.DB.Transaction(func(tx *gorm.DB) error{
		for _, ri := range cart.Items{
			id64, err := strconv.ParseUint(ri.SKU, 10, 32)
			if err != nil {
				return fmt.Errorf("invalid sku ID %q: %w", ri.SKU, err)
			}
			skuItemID:= uint(id64)

			if err := tx.Model(&model.SKUItem{}).
			Where("ID=?", skuItemID).
			UpdateColumn("quantity", gorm.Expr("quantity - ?", ri.Quantity)).
			Error; err != nil{
				return fmt.Errorf("updating SKUItem %d quantity: %w", skuItemID, err)
			}

			if err := tx.Model(&model.SKU{}).
			Where("ID=?", skuItemID).
			UpdateColumn("pick_count", gorm.Expr("pick_count + ?", 1)).
			Error; err != nil{
				return fmt.Errorf("updating SKUItem %d pick count: %w", skuItemID, err)
			}
		}
		return nil
	})
}