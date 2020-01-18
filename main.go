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

func SortBy(homes []home, less func(a, b home) bool) []home{
	result := make([]home, len(homes))
	copy(result, homes)
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}

func SortByPriceAsc(homes []home) []home {
	return SortBy(homes, func(a, b home) bool {
		if a.price < b.price {
			return true
		}
		return false
	})
}

func SortByPriceDesc(homes []home) []home {
	return SortBy(homes, func(a, b home) bool {
		if a.price > b.price{
			return true
		}
		return false
	})
}

func SortByDistanceFromCenterAsc(homes []home) []home {
	return SortBy(homes, func(a, b home) bool {
		return a.distanceFromCenter < b.distanceFromCenter
	})
}

func SortByDistanceFromCenterDesc(homes []home) []home {
	return SortBy(homes, func(a, b home) bool {
		return a.distanceFromCenter > b.distanceFromCenter
	})
}

func FilterBy(homes []home, less func(a home) bool) []home{
	result := make([]home, 0)
	for _, home := range homes{
		if less(home){
			result = append(result, home)
		}
	}
	return result
}

func SearchByMaxPrice(homes []home, maxPrice int64) []home{
	return FilterBy(homes, func(a home) bool {
		if a.price <= maxPrice{
			return true
		}
		return false
	})
}

func SearchFromXPriceToYPrice(homes []home, minPrice, maxPrice int64) []home{
	return FilterBy(homes, func(a home) bool {
		if a.price >= minPrice && a.price <= maxPrice{
			return true
		}
		return false
	})
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
