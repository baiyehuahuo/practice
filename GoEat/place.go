package main

import (
	"fmt"
	"strconv"

	"github.com/liushuochen/gotable"
)

type Place struct {
	name       string
	ress       []*Restaurant
	drinksMenu []*Product
}

func (pls *Place) GetWeight() float64 {
	ans := 0.0
	for i := range pls.ress {
		ans += pls.ress[i].GetWeight()
	}
	return ans
}

func (pls *Place) getTotalPrice() (ans float64) {
	for i := range pls.ress {
		ans += pls.ress[i].getTotalPrice()
	}
	return ans
}

func (pls *Place) Print() {
	drinkPrice := getAverageDrinkPrice(pls.drinksMenu)
	fmt.Printf("Average Drink Price: %f\n", drinkPrice)
	table, err := gotable.Create("Restaurant", "Choose", "Price", "OfferDrink", "Total Price")
	if err != nil {
		panic(err)
	}
	values := make(map[string]string)
	for _, res := range pls.ress {
		values["Restaurant"] = res.name
		for _, pro := range res.menu {
			values["Choose"] = pro.name
			values["Price"] = strconv.FormatFloat(pro.price, 'f', 1, 64)
			values["OfferDrink"] = strconv.FormatBool(pro.offerDrink)
			if !pro.offerDrink {
				values["Total Price"] = strconv.FormatFloat(pro.price+drinkPrice, 'f', 1, 64)
			} else {
				values["Total Price"] = strconv.FormatFloat(pro.price, 'f', 1, 64)
			}
			if err = table.AddRow(values); err != nil {
				panic(err)
			}
		}
	}
	table.CloseBorder()
	fmt.Printf("%v", table)
	fmt.Printf("Average weight price: %v\n\n", pls.getTotalPrice()/pls.GetWeight())
}
