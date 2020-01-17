package main

import (
	"fmt"
)

var homes = []home{
	{1, "Двухкомнатная квартира", 300_000, 30, "Sino"},
	{2, "Двухкомнатная квартира", 50_000,20, "Firdavsi"},
	{3, "3-комнатная квартира", 35_000, 10, "Sino"},
	{4, "4-мнатная квартира", 1_000_000, 1, "Sino"},
	{5, "Двухкомнатная квартира", 300_000, 30, "Firdavsi"},
	{6, "1-комнатная квартира", 150_000, 25, "Shohmansur"},
	{7, "8-комнатная квартира", 3_500_000, 1, "Somoni"},
}

func ExampleSortByPriceAsc() {
	result := SortByPriceAsc(homes)
	fmt.Println(result)
	// Output: [{3 3-комнатная квартира 35000 10 Sino} {2 Двухкомнатная квартира 50000 20 Firdavsi} {6 1-комнатная квартира 150000 25 Shohmansur} {1 Двухкомнатная квартира 300000 30 Sino} {5 Двухкомнатная квартира 300000 30 Firdavsi} {4 4-мнатная квартира 1000000 1 Sino} {7 8-комнатная квартира 3500000 1 Somoni}]
}

func ExampleSortByPriceDesc() {
	result := SortByPriceDesc(homes)
	fmt.Print(result)
	// Output: [{7 8-комнатная квартира 3500000 1 Somoni} {4 4-мнатная квартира 1000000 1 Sino} {5 Двухкомнатная квартира 300000 30 Firdavsi} {1 Двухкомнатная квартира 300000 30 Sino} {6 1-комнатная квартира 150000 25 Shohmansur} {2 Двухкомнатная квартира 50000 20 Firdavsi} {3 3-комнатная квартира 35000 10 Sino}]
}

func ExampleSortByDistanceFromCenterAsc()  {
	result := SortByDistanceFromCenterAsc(homes)
	fmt.Print(result)
	// Output: [{7 8-комнатная квартира 3500000 1 Somoni} {4 4-мнатная квартира 1000000 1 Sino} {3 3-комнатная квартира 35000 10 Sino} {2 Двухкомнатная квартира 50000 20 Firdavsi} {6 1-комнатная квартира 150000 25 Shohmansur} {5 Двухкомнатная квартира 300000 30 Firdavsi} {1 Двухкомнатная квартира 300000 30 Sino}]
}

func ExampleSortByDistanceFromCenterDesc()  {
	result := SortByDistanceFromCenterDesc(homes)
	fmt.Print(result)
	// Output: [{1 Двухкомнатная квартира 300000 30 Sino} {5 Двухкомнатная квартира 300000 30 Firdavsi} {6 1-комнатная квартира 150000 25 Shohmansur} {2 Двухкомнатная квартира 50000 20 Firdavsi} {3 3-комнатная квартира 35000 10 Sino} {4 4-мнатная квартира 1000000 1 Sino} {7 8-комнатная квартира 3500000 1 Somoni}]
}

var limPrice int64 = 1_000_000
func ExampleSearchByMaxPrice()  {
	result := SearchByMaxPrice(homes, limPrice)
	fmt.Print(result)
	// Output: [{1 Двухкомнатная квартира 300000 30 Sino} {2 Двухкомнатная квартира 50000 20 Firdavsi} {3 3-комнатная квартира 35000 10 Sino} {4 4-мнатная квартира 1000000 1 Sino} {5 Двухкомнатная квартира 300000 30 Firdavsi} {6 1-комнатная квартира 150000 25 Shohmansur}]
}

var maxPrice int64 = 1_000_000
var minPrice int64 = 0
func ExampleSearchFromXPriceToYPrice()  {
	result := SearchFromXPriceToYPrice(homes, minPrice, maxPrice)
	fmt.Print(result)
	// Output: [{1 Двухкомнатная квартира 300000 30 Sino} {2 Двухкомнатная квартира 50000 20 Firdavsi} {3 3-комнатная квартира 35000 10 Sino} {4 4-мнатная квартира 1000000 1 Sino} {5 Двухкомнатная квартира 300000 30 Firdavsi} {6 1-комнатная квартира 150000 25 Shohmansur}]
}

var regions = []string{"Sino", "Firdavsi"}
func ExampleSearchByRegion()  {
	result := SearchByRegion(homes, regions)
	fmt.Print(result)
	// Output: [{1 Двухкомнатная квартира 300000 30 Sino} {2 Двухкомнатная квартира 50000 20 Firdavsi} {3 3-комнатная квартира 35000 10 Sino} {4 4-мнатная квартира 1000000 1 Sino} {5 Двухкомнатная квартира 300000 30 Firdavsi}]
}