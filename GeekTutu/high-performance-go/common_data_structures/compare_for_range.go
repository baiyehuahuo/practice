package common_data_structures

type Item struct {
	id  int
	val [4096]byte
}

func generateItems(n int) []*Item {
	items := make([]*Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, &Item{id: i})
	}
	return items
}
