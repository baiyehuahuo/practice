package main

func Xiaowai() *Place {
	var drinksMenu []*Product
	drinksMenu = append(drinksMenu, &Product{name: "百世", weight: 1, price: 2})
	drinksMenu = append(drinksMenu, &Product{name: "可口", weight: 1, price: 2})
	drinksMenu = append(drinksMenu, &Product{name: "芬达", weight: 1, price: 2})
	drinksMenu = append(drinksMenu, &Product{name: "雪碧", weight: 1, price: 2})
	drinksMenu = append(drinksMenu, &Product{name: "水", weight: 1, price: 0})

	var ress []*Restaurant
	var res *Restaurant

	res = &Restaurant{name: "汉堡王"}
	res.menu = append(res.menu, &Product{name: "四件套", weight: 1, price: 25, offerDrink: true})
	ress = append(ress, res)

	res = &Restaurant{name: "板烧厨房"}
	res.menu = append(res.menu, &Product{name: "板烧+汉堡肉", weight: 2, price: 25, offerDrink: true})
	ress = append(ress, res)

	res = &Restaurant{name: "面霸"}
	res.menu = append(res.menu, &Product{name: "杂酱拌面", weight: 4, price: 13, drinkMenu: drinksMenu})
	res.menu = append(res.menu, &Product{name: "土豆牛肉拌面", weight: 1, price: 16, drinkMenu: drinksMenu})
	ress = append(ress, res)

	return &Place{
		name:       "校外",
		ress:       ress,
		drinksMenu: drinksMenu,
	}
}
