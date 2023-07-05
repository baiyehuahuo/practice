package main

import (
	"fmt"
	"math/rand"
)

type Weight interface {
	GetWeight() float64
}

var _ Weight = &Place{}
var _ Weight = &Product{}
var _ []Weight = append([]Weight{}, &Product{})

func RandomSelect(ws []Weight) interface{} {
	totalWeights := 0.0
	for i := range ws {
		totalWeights += ws[i].GetWeight()
	}
	idxWeight := rand.Float64() * totalWeights
	for i := range ws {
		w := ws[i].GetWeight()
		if w > idxWeight {
			return ws[i]
		}
		idxWeight -= w
	}
	return nil
}

type Restaurant struct {
	name string
	menu []*Product
}

func (res *Restaurant) GetWeight() float64 {
	ans := 0.0
	for i := range res.menu {
		ans += res.menu[i].weight
	}
	return ans
}

func (res *Restaurant) ToString() string {
	return ""
}

func (res *Restaurant) getTotalPrice() float64 {
	ans := 0.0
	for i := range res.menu {
		ans += res.menu[i].weight * res.menu[i].price
		if !res.menu[i].offerDrink {
			ans += getAverageDrinkPrice(res.menu[i].drinkMenu)
		}
	}
	return ans
}

type Product struct {
	name       string
	weight     float64
	price      float64
	offerDrink bool
	drinkMenu  []*Product
}

func getAverageDrinkPrice(drinksMenu []*Product) float64 {
	totalDrinkWeights, totalDrinkPrice := 0.0, 0.0
	for _, drink := range drinksMenu {
		totalDrinkWeights += drink.weight
		totalDrinkPrice += drink.weight * drink.price
	}
	return totalDrinkPrice / totalDrinkWeights
}

func (pro *Product) GetWeight() float64 {
	return pro.weight
}

func (pro *Product) ToString() string {
	return fmt.Sprintf("菜名: %s 价格: %.2f", pro.name, pro.price)
}
