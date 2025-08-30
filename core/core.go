package core

type TeaMaster struct {
}

type GeneralStats struct {
	TotalWeight  float32
	TotalPrice   float32
	NumberOfTeas int
	GramsDrunk   float32
}

func (t TeaMaster) AddTea(name string, year int, stock float32, price float32, tags *[]string) {
	tea := &Tea{
		Name:         name,
		Year:         year,
		Stock:        stock,
		PricePerGram: price,
	}

	saveTea(tea)

	if len(*tags) > 0 {
		assignTags(tea, tags)
	}
}

func (t TeaMaster) GetGenetalStats(teaList []Tea) *GeneralStats {
	var weight float32 = 0.00
	var price float32 = 0.00

	for _, tea := range teaList {
		weight += tea.Stock
		price += tea.Stock * tea.PricePerGram
	}

	stats := &GeneralStats{
		NumberOfTeas: len(teaList),
		TotalWeight:  weight,
		TotalPrice:   price,
	}

	return stats
}

func (t TeaMaster) GetTeasWithFilters(filters *[][]string, sorting *[]string) []Tea {
	return GetTeasWithFilters(*filters, *sorting)
}

func (t TeaMaster) GetAllTeas() []Tea {
	return getAllTeas()
}

func (t TeaMaster) DeleteTea(id int) {
	deleteTea(id)
}

func InitCore() *TeaMaster {
	init_db()

	return &TeaMaster{}
}
