package main

import (
	"log"
	"github.com/stilln0thing/WareManager/server/internal/db"
	"github.com/stilln0thing/WareManager/server/internal/model"
	"github.com/stilln0thing/WareManager/server/internal/service"

	// "github.com/stilln0thing/WareManager/server/internal/seeders"
	// "time"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatal("Failed to initialise the database")
	}

	db.DB.AutoMigrate(&model.SKU{})
	db.DB.AutoMigrate(&model.SKUItem{})
	// for _, sku := range seeders.SeedSKUs {
    //     db.DB.FirstOrCreate(&model.SKU{}, model.SKU{ID: sku.ID, Name: sku.Name, Category: sku.Category, CreatedAt: time.Now()})
    // }
	// for _, skuItems := range seeders.SeedSKUItems {
    //     db.DB.FirstOrCreate(&model.SKUItem{}, model.SKUItem{ID: skuItems.ID,SKUId: skuItems.SKUId,InTime: skuItems.InTime,
	// 	AisleNo: skuItems.AisleNo, Quantity: skuItems.Quantity, CreatedAt: time.Now()})
    // }
	router := gin.Default()
	router.POST("/order",func(c *gin.Context) {
    	if err := service.HandleOrder(); err != nil {
       		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        	return
    	}
		c.JSON(http.StatusOK, gin.H{
        	"status":  "success",
        	"message": "Order processed successfully",
    })

	})
		log.Println("Server is running on :8080")
    	if err := router.Run(":8081"); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}


// i have to make a recommendation model on which item to place where based on the pick count so i will take a generalisation that 
// the closest aisle would be the one with the smallest number 