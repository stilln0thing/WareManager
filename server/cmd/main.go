package main

import (
	"github.com/stilln0thing/WareManager/server/internal/db"
	"github.com/stilln0thing/WareManager/server/internal/model"
)

func main() {
	db.Init()

	db.DB.AutoMigrate(&model.SKU{})
}
