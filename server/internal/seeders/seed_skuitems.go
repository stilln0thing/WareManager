package seeders

import (
    "time"
	"github.com/stilln0thing/WareManager/server/internal/model"

)

var SeedSKUItems = []model.SKUItem{}

func init() {
    now := time.Now().Truncate(time.Second)
    for id := uint(1); id <= 50; id++ {
        SeedSKUItems = append(SeedSKUItems, model.SKUItem{
            ID:        id,
            SKUId:     id,       
            InTime:    now,      
            AisleNo:   int(id%10) + 1, // sample 10 aisles for now
            Quantity:  100,    
            CreatedAt: now,
        })
    }
}