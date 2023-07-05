package main

func Xiyuan() *Place {
	var ress []*Restaurant
	var drinksMenu []*Product
	drinksMenu = append(drinksMenu, &Product{name: "百世", weight: 1, price: 1.9})
	drinksMenu = append(drinksMenu, &Product{name: "可口", weight: 1, price: 1.9})
	drinksMenu = append(drinksMenu, &Product{name: "芬达", weight: 1, price: 1.9})
	drinksMenu = append(drinksMenu, &Product{name: "雪碧", weight: 1, price: 1.9})
	drinksMenu = append(drinksMenu, &Product{name: "菠萝啤", weight: 4, price: 0.9})
	drinksMenu = append(drinksMenu, &Product{name: "水", weight: 14, price: 1.0})

	res := &Restaurant{name: "乡村餐厅"}
	res.menu = append(res.menu, &Product{name: "自选", weight: 33, price: 10, drinkMenu: drinksMenu})
	res.menu = append(res.menu, &Product{name: "火腿炒蛋", weight: 2, price: 13, drinkMenu: drinksMenu})
	res.menu = append(res.menu, &Product{name: "蒜苗炒肉", weight: 1, price: 16, drinkMenu: drinksMenu})
	ress = append(ress, res)

	res = &Restaurant{name: "广州正宗肠粉"}
	res.menu = append(res.menu, &Product{name: "炒饭", weight: 7, price: 11, drinkMenu: drinksMenu})
	res.menu = append(res.menu, &Product{name: "炒面", weight: 1, price: 11, drinkMenu: drinksMenu})
	ress = append(ress, res)

	res = &Restaurant{name: "黄鼠狼鸡公煲"}
	res.menu = append(res.menu, &Product{name: "方便面", weight: 2, price: 21, drinkMenu: drinksMenu})
	ress = append(ress, res)

	res = &Restaurant{name: "鲍汁黄焖鸡"}
	res.menu = append(res.menu, &Product{name: "香酥肉末粉丝", weight: 4, price: 13, offerDrink: true})
	res.menu = append(res.menu, &Product{name: "土豆牛肉", weight: 1, price: 15, offerDrink: true})
	ress = append(ress, res)

	res = &Restaurant{name: "排骨饭"}
	res.menu = append(res.menu, &Product{name: "小份酱汁排骨饭", weight: 1, price: 15, drinkMenu: drinksMenu})
	ress = append(ress, res)

	res = &Restaurant{name: "清真餐厅"}
	res.menu = append(res.menu, &Product{name: "牛肉拉面", weight: 1, price: 10, offerDrink: true})
	ress = append(ress, res)

	res = &Restaurant{name: "福建千里香馄饨"}
	res.menu = append(res.menu, &Product{name: "馄饨+包子", weight: 1, price: 7 + 7, offerDrink: true})
	ress = append(ress, res)

	return &Place{
		name:       "西苑",
		ress:       ress,
		drinksMenu: drinksMenu,
	}
}
