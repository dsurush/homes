package main

import (
	"sort"
)

type home struct {
	id int64
	name string
	price int64
	distanceFromCenter int64
	region string
}

func SortByPriceAsc(homes []home) []home {
	result := make([]home, len(homes))
	copy(result, homes)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price < result[j].price
	})
	return result
}

func SortByPriceDesc(homes []home) []home {
	result := make([]home, len(homes))
	copy(result, homes)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price > result[j].price
	})
	return result
}

func SortByDistanceFromCenterAsc(homes []home) []home {
	result := make([]home, len(homes))
	copy(result, homes)
	sort.Slice(result, func(i, j int) bool {
		return result[i].distanceFromCenter < result[j].distanceFromCenter
	})
	return result
}

func SortByDistanceFromCenterDesc(homes []home) []home {
	result := make([]home, len(homes))
	copy(result, homes)
	sort.Slice(result, func(i, j int) bool {
		return result[i].distanceFromCenter > result[j].distanceFromCenter
	})
	return result
}

func SearchByMaxPrice(homes []home, maxPrice int64) []home{
	result := make([]home, 0)
	for _, home := range homes{
		if home.price <= maxPrice{
			result = append(result, home)
		}
	}
	return result
}

func SearchFromXPriceToYPrice(homes []home, minPrice, maxPrice int64) []home{
	result := make([]home, 0)
	for _, home := range homes{
		if home.price >= minPrice && home.price <= maxPrice{
			result = append(result, home)
		}
	}
	return result
}

func SearchByRegion(homes []home, regions []string) []home{
	result := make([]home, 0)
	for _, home := range homes{
		for _, region := range regions{
			if home.region == region{
				result = append(result, home)
				break
			}
		}
	}
	return result
}

func main() {}
