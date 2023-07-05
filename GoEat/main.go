package main

import (
	"fmt"
)

func main() {
	pls := []*Place{Xiyuan(), Houjie(), Xiaowai()}
	totalWeights, totalPrice := 0.0, 0.0
	for i := range pls {
		weight, price := pls[i].GetWeight(), pls[i].getTotalPrice()
		totalWeights += weight
		totalPrice += price
		pls[i].Print()
	}
	fmt.Println(totalPrice / float64(totalWeights))
	pl := RandomSelect(PlacesToWeights(pls)).(*Place)
	res := RandomSelect(RessToWeights(pl.ress)).(*Restaurant)
	pro := RandomSelect(ProsToWeights(res.menu)).(*Product)
	totPrice := pro.price
	fmt.Printf("地点: %s\n餐馆: %s\n%s\n", pl.name, res.name, pro.ToString())
	if !pro.offerDrink {

		drink := RandomSelect(ProsToWeights(pro.drinkMenu)).(*Product)
		fmt.Printf("饮料: %s 价格: %.2f\n", drink.name, drink.price)
		totPrice += drink.price
	}
	fmt.Printf("总价格: %.2f\n", totPrice)
}

func PlacesToWeights(ws []*Place) []Weight {
	weights := make([]Weight, len(ws))
	for i := range ws {
		weights[i] = ws[i]
	}
	return weights
}

func RessToWeights(ws []*Restaurant) []Weight {
	weights := make([]Weight, len(ws))
	for i := range ws {
		weights[i] = ws[i]
	}
	return weights
}

func ProsToWeights(ws []*Product) []Weight {
	weights := make([]Weight, len(ws))
	for i := range ws {
		weights[i] = ws[i]
	}
	return weights
}
