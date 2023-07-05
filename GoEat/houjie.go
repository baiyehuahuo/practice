package main

func Houjie() *Place {
	var drinksMenu []*Product
	drinksMenu = append(drinksMenu,
		&Product{name: "百世", weight: 1, price: 2},
		&Product{name: "可口", weight: 1, price: 2},
		&Product{name: "芬达", weight: 1, price: 2},
		&Product{name: "雪碧", weight: 1, price: 2},
		&Product{name: "矮罐菠萝啤", weight: 1, price: 1.2},
		&Product{name: "高罐菠萝啤", weight: 1, price: 2},
		&Product{name: "水", weight: 6, price: 1.2},
	)

	var ress []*Restaurant
	res := &Restaurant{name: "黄焖鸡米饭"}
	res.menu = append(res.menu, &Product{name: "小份黄焖鸡", weight: 5, price: 16, drinkMenu: drinksMenu})
	res.menu = append(res.menu, &Product{name: "排骨饭", weight: 3, price: 20, drinkMenu: drinksMenu})
	ress = append(ress, res)

	res = &Restaurant{name: "麻辣烫"}
	res.menu = append(res.menu, &Product{name: "麻辣烫", weight: 7, price: 14, drinkMenu: drinksMenu})
	ress = append(ress, res)

	res = &Restaurant{name: "李厨"}
	res.menu = append(res.menu, &Product{name: "四季豆炒肉", weight: 3, price: 14, drinkMenu: drinksMenu})
	res.menu = append(res.menu, &Product{name: "土豆丝炒肉", weight: 4, price: 13, drinkMenu: drinksMenu})
	ress = append(ress, res)

	res = &Restaurant{name: "小哥炒饭"}
	res.menu = append(res.menu, &Product{name: "老干妈火腿炒饭", weight: 7, price: 12, drinkMenu: drinksMenu})
	res.menu = append(res.menu, &Product{name: "火腿炒面", weight: 1, price: 11, drinkMenu: drinksMenu})
	ress = append(ress, res)

	res = &Restaurant{name: "重庆面馆"}
	res.menu = append(res.menu, &Product{name: "大份牛肉面", weight: 2, price: 12, drinkMenu: drinksMenu})
	ress = append(ress, res)

	res = &Restaurant{name: "筒子骨面粉"}
	res.menu = append(res.menu, &Product{name: "四季豆炒肉", weight: 4, price: 13, drinkMenu: drinksMenu})
	res.menu = append(res.menu, &Product{name: "牛肉炒饭", weight: 3, price: 13, drinkMenu: drinksMenu})
	ress = append(ress, res)

	return &Place{
		name:       "后街",
		ress:       ress,
		drinksMenu: drinksMenu,
	}
}
