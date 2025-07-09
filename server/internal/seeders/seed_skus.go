package seeders

import (
    "github.com/stilln0thing/WareManager/server/internal/model"
)

var SeedSKUs = []model.SKU{
    {ID: 1,  Name: "Banana",                         Category: "Fruits"},
    {ID: 2,  Name: "Apple",                          Category: "Fruits"},
    {ID: 3,  Name: "Orange",                         Category: "Fruits"},
    {ID: 4,  Name: "Grapes",                         Category: "Fruits"},
    {ID: 5,  Name: "Tomato",                         Category: "Vegetables"},
    {ID: 6,  Name: "Cucumber",                       Category: "Vegetables"},
    {ID: 7,  Name: "Carrot",                         Category: "Vegetables"},
    {ID: 8,  Name: "Potato",                         Category: "Vegetables"},
    {ID: 9,  Name: "Onion",                          Category: "Vegetables"},
    {ID: 10, Name: "Green Chillies",                 Category: "Vegetables"},

    {ID: 11, Name: "Bread Loaf",                     Category: "Bakery"},
    {ID: 12, Name: "Milk 1L",                        Category: "Dairy"},
    {ID: 13, Name: "Curd 500g",                      Category: "Dairy"},
    {ID: 14, Name: "Butter 200g",                    Category: "Dairy"},
    {ID: 15, Name: "Cheese Slices",                  Category: "Dairy"},
    {ID: 16, Name: "Eggs (12pcs)",                   Category: "Dairy"},
    {ID: 17, Name: "Paneer 200g",                    Category: "Dairy"},
    {ID: 18, Name: "Rice 1kg",                       Category: "Staples"},
    {ID: 19, Name: "Wheat Flour 1kg",                Category: "Staples"},
    {ID: 20, Name: "Sugar 1kg",                      Category: "Staples"},
    {ID: 21, Name: "Salt 1kg",                       Category: "Staples"},
    {ID: 22, Name: "Tea 250g",                       Category: "Beverages"},
    {ID: 23, Name: "Coffee 200g",                    Category: "Beverages"},
    {ID: 24, Name: "Vegetable Oil 1L",               Category: "Staples"},
    {ID: 25, Name: "Daal (Toor) 1kg",                Category: "Staples"},

    {ID: 26, Name: "Chips (Kurkure)",                Category: "Snacks"},
    {ID: 27, Name: "Namkeen (Sev Mamra)",            Category: "Snacks"},
    {ID: 28, Name: "Cookies Pack",                   Category: "Snacks"},
    {ID: 29, Name: "Biscuits Pack",                  Category: "Snacks"},
    {ID: 30, Name: "Ice Cream (500ml)",              Category: "Frozen"},
    {ID: 31, Name: "Juice Bottle 1L",                Category: "Beverages"},
    {ID: 32, Name: "Soft Drink Can",                 Category: "Beverages"},
    {ID: 33, Name: "Mineral Water 1L",               Category: "Beverages"},

    {ID: 34, Name: "Shampoo 200ml",                  Category: "Personal Care"},
    {ID: 35, Name: "Soap Bar",                       Category: "Personal Care"},
    {ID: 36, Name: "Toothpaste 100g",                Category: "Personal Care"},
    {ID: 37, Name: "Toothbrush",                     Category: "Personal Care"},
    {ID: 38, Name: "Detergent 500g",                 Category: "Household"},
    {ID: 39, Name: "Dishwash Bar",                   Category: "Household"},
    {ID: 40, Name: "Sanitary Pads Pack",             Category: "Personal Care"},
    {ID: 41, Name: "Diapers Pack",                   Category: "Baby Care"},

    {ID: 42, Name: "Cold Drink 2L",                  Category: "Beverages"},
    {ID: 43, Name: "Chicken Curry Cut 500g",         Category: "Meat"},
    {ID: 44, Name: "Fish Fillet 500g",               Category: "Meat"},
    {ID: 45, Name: "Ready-to-Eat Curry Pack",        Category: "Ready Meals"},
    {ID: 46, Name: "Frozen Paratha (10pcs)",         Category: "Frozen"},
    {ID: 47, Name: "Frozen Samosa (6pcs)",           Category: "Frozen"},
    {ID: 48, Name: "Frozen Dosa Batter 1L",          Category: "Frozen"},
    {ID: 49, Name: "Paneer Butter Masala (ready)",   Category: "Ready Meals"},
    {ID: 50, Name: "Instant Noodles Pack",           Category: "Instant Foods"},
}
